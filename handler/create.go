package handler

import (
	"errors"
	"net/http"
	"net/mail"
	"regexp"

	"example.com/morethanjustlinks/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	CREATE_USERS_TABLE = `CREATE TABLE users (
		uuid VARCHAR(36) NOT NULL,
		name VARCHAR(128) NOT NULL,
		email VARCHAR(255) NOT NULL,
		phone VARCHAR(10) NOT NULL,
		psword VARCHAR(255) NOT NULL,
		verified BOOLEAN,
		PRIMARY KEY (uuid)
	);
	`

	CREATE_LINKS_TABLE = `CREATE TABLE links (
		username VARCHAR(128) NOT NULL,
		uuid VARCHAR(36) NOT NULL,
		name VARCHAR(128) NOT NULL,
		url VARCHAR(255) NOT NULL,
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

type HasherInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

// SetupService is used to create tables required for service
// TODO this should happen at init, but for demomstatstration purposes it is here
// TODO add checks before just dropping tables
func (h *HandlerService) SetupService(ctx *gin.Context) {
	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	h.DropUsersTable(ctx)
	h.DropLinksTable(ctx)

	if len(ctx.Errors) == 0 {
		h.CreateUsersTable(ctx)
		h.CreateLinksTable(ctx)
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"msg": "tables are successfully created"})

}

func (h *HandlerService) CreateUsersTable(ctx *gin.Context) {
	_, err := h.maria_repo.Exec(CREATE_USERS_TABLE)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error dropping links"})
		ctx.Error(err)
	}
	ctx.Next()

}

func (h *HandlerService) CreateLinksTable(ctx *gin.Context) {
	_, err := h.maria_repo.Exec(CREATE_LINKS_TABLE)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error dropping links"})
		ctx.Error(err)
	}
	ctx.Next()
}

// CreateLink is used to create links for users
// TODO validate the input data
func (h *HandlerService) CreateLink(ctx *gin.Context) {

	// todo bind request data

	_, err := h.maria_repo.Exec("insert into links (username,uuid,name,url) values (?,?,?,?);",
		"ham", "30a1ce10-e885-4652-a9cc-8c2bff55f8f2", "morethanjustlinks", "morethanjustlinks.com")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})

	}
}

// NewAccount is used to create a new user accounts
// TODO send verification email for accounts created before allowing edits
func (h *HandlerService) NewAccount(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	// bind and validate input data
	req, err := validateInsertUserRequest(ctx)
	if err != nil {
		h.sugaredLogger.Errorw("error validating insert user request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please enter the required request fields"})
		return
	}

	rows, err := h.maria_repo.Query("SELECT COUNT(*) FROM users WHERE name = ?", req.Name)

	defer func() {
		h.sugaredLogger.Desugar().Sync()
		// rows.Close()
	}()

	if err != nil {
		h.sugaredLogger.Errorw("error counting number of usersnames", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	userAlreadyExists, _ := UserAlreadyExists(rows)
	if userAlreadyExists {
		h.sugaredLogger.Errorw("error checking user already exists", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	// encrypt password
	hashPswd, err := HashPassword(req.Psword)
	if err != nil {
		h.sugaredLogger.Errorw("error hashing and salting", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	// insert into db
	result, err := h.maria_repo.Exec(
		"INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)",
		uuid.New().String(), req.Name, req.Email, req.Phone, hashPswd, req.Verified,
	)

	if err != nil {
		h.sugaredLogger.Errorw("error inserting users into db", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	_, err = result.LastInsertId()
	if err != nil {
		h.sugaredLogger.Errorw("Error retrieving last uuid", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "successfully created new user",
	})
}

func validateInsertUserRequest(ctx *gin.Context) (InsertUserRequest, error) {
	var req InsertUserRequest

	if err := ctx.BindJSON((&req)); err != nil {
		return req, err
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

func UserAlreadyExists(rows db.RowsInterface) (bool, error) {
	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false, err
		}

		if count > 0 {
			return true, nil
		}

	}
	return false, nil
}
