package handler

import (
	"errors"
	"net/http"
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	CREATE_USER_TABLE = `CREATE TABLE users (
		uuid VARCHAR(36) NOT NULL,
		name VARCHAR(128) NOT NULL,
		email VARCHAR(255) NOT NULL,
		phone VARCHAR(10) NOT NULL,
		psword VARCHAR(255) NOT NULL,
		verified BOOLEAN,
		PRIMARY KEY (uuid)
	);
	`
)

type InsertUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Psword   string `json:"psword"`
	Verified bool   `json:"verified,omitempty"`
}

// SetupService is used to create tables required for service
func (h *HandlerService) SetupService(ctx *gin.Context) {
	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	h.DropUsersTable(ctx)

	if len(ctx.Errors) == 0 {
		h.CreateUserTable(ctx)
	}

}

func (h *HandlerService) CreateUserTable(ctx *gin.Context) {
	_, err := h.maria_repo.Exec(CREATE_USER_TABLE)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating users"})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"msg": "created user tables succesfully"})

}

func (h *HandlerService) InsertUser(ctx *gin.Context) {
	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	ctx.Header("Content-Type", "application/json")

	// bind and validate input data
	req, err := validateInsertUserRequest(ctx)
	if err != nil {
		h.sugaredLogger.Errorw("Error with input data", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error with input data"})
		return
	}

	// encrypt password
	hashPswd, err := HashPassword(req.Psword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong with your login request...01"})
		return
	}

	// insert into db
	result, err := h.maria_repo.Exec(
		"INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)",
		uuid.New().String(), req.Name, req.Email, req.Phone, hashPswd, req.Verified,
	)

	if err != nil {
		h.sugaredLogger.Errorw("Error adding new user", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding new user"})
		return
	}

	_, err = result.LastInsertId()
	if err != nil {
		h.sugaredLogger.Errorw("Error retrieving last uuid", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching last id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "successfully created new user",
		"req": req,
	})
}

func validateInsertUserRequest(ctx *gin.Context) (InsertUserRequest, error) {
	var req InsertUserRequest

	if err := ctx.BindJSON((&req)); err != nil {
		return req, errors.New("please enter a proper request to insert new user")
	}

	if req.Name == "" || req.Email == "" || req.Phone == "" || req.Psword == "" {
		return req, errors.New("please enter the required request fields")
	}

	if !isValidUsername(req.Name) {
		return req, errors.New("please enter a valid username")
	}

	if !isValidEmail(req.Email) {
		return req, errors.New("please enter a valid email")
	}

	if !isValidPhoneNumber(req.Phone) {
		return req, errors.New("please enter a valid phone")
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
