package handler

import (
	"errors"
	"net/http"
	"net/mail"
	"regexp"

	"example.com/morethanjustlinks/models"
	"example.com/morethanjustlinks/presentation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type InsertUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Psword   string `json:"psword"`
	Verified bool   `json:"verified,omitempty"`
	Role     string `json:"role,omitempty"`
}

type HasherInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

// NewAccount is used to create a new user accounts
// TODO send verification email for accounts created before allowing edits
func (h *Handler) NewAccount(ctx *gin.Context) {
	var user models.User
	var err error

	req, err := validateInsertUserRequest(ctx)
	if err != nil {
		h.sugaredLogger.Errorw("error validating insert user request", zap.Any("req", req), zap.Any("err", err.Error()))
		ctx.JSON(http.StatusBadRequest, handleValidationErrors(err))
		return
	}

	err = h.db.Where("phone = ? OR email = ?", req.Phone, req.Email).First(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		h.sugaredLogger.Warnw("expected record not found",
			zap.String("email", req.Email),
			zap.String("phone", req.Phone),
		)
		ctx.JSON(http.StatusOK, gin.H{"error": "email or phone number entered is already in use"})
		return
	}

	// encrypt password
	hashPswd, err := HashPassword(req.Psword)
	if err != nil {
		h.sugaredLogger.Errorw("error hashing and salting", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	// save the user to db
	user = models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Psword:   hashPswd,
		Role:     req.Role,
		Verified: req.Verified,
	}
	result := h.db.Create(&user)

	if result.Error != nil {
		h.sugaredLogger.Errorw("error", zap.String("username", req.Name))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	sendVerificationEmail(user)

	ctx.JSON(http.StatusOK, gin.H{"message": "Thanks for joining! You please check your email for a verification link."})
}

// sendVerificationEmail TODO implement verification email
func sendVerificationEmail(user models.User) {
	if user.Email == "" {
		return
	}
}

func handleValidationErrors(err error) gin.H {
	switch err.(type) {
	case *presentation.ErrInvalidUsername:
		errors := map[string]string{"name": err.Error()}
		return gin.H{"errors": errors}
	case *presentation.ErrInvalidEmail:
		errors := map[string]string{"email": err.Error()}
		return gin.H{"errors": errors}
	case *presentation.ErrInvalidPhone:
		errors := map[string]string{"phone": err.Error()}
		return gin.H{"errors": errors}

	}
	return gin.H{}
}

func validateInsertUserRequest(ctx *gin.Context) (models.User, error) {
	var req models.User

	if err := ctx.BindJSON((&req)); err != nil {
		return req, err
	}

	if req.Name == "" || req.Email == "" || req.Phone == "" || req.Psword == "" {
		return req, errors.New("please enter the required request fields")
	}

	if !isValidUsername(req.Name) {
		return req, &presentation.ErrInvalidUsername{}
	}

	if !isValidEmail(req.Email) {
		return req, &presentation.ErrInvalidEmail{}
	}

	if !isValidPhoneNumber(req.Phone) {
		return req, &presentation.ErrInvalidPhone{}
	}

	return req, nil
}

func validateUpdateUserRequest(ctx *gin.Context) (models.User, error) {
	var req models.User

	if err := ctx.BindJSON((&req)); err != nil {
		return req, err
	}

	if req.Name == "" || req.Email == "" || req.Phone == "" {
		return req, errors.New("please enter the required request fields")
	}

	if !isValidUsername(req.Name) {
		return req, &presentation.ErrInvalidUsername{}
	}

	if !isValidEmail(req.Email) {
		return req, &presentation.ErrInvalidEmail{}
	}

	if !isValidPhoneNumber(req.Phone) {
		return req, &presentation.ErrInvalidPhone{}
	}

	return req, nil
}

func isValidUsername(name string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z_]{3,25}$", name)
	return match
}

// this accepts some@email without .com should this be accepted?
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidPhoneNumber(phone string) bool {
	match, _ := regexp.MatchString("^[0-9]{10}$", phone)
	return match
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
