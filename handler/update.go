package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateUserRequest struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Psword   string `json:"psword"`
	Verified bool   `json:"verified,omitempty"`
}

func (h HandlerService) UpdateUser(ctx *gin.Context) {
	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	ctx.Header("Content-Type", "application/json")

	// bind and validate input data
	req, err := validateUpdateUserRequest(ctx)
	if err != nil {
		h.sugaredLogger.Errorw("Error with update data", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error with update data"})
		return
	}

	// TODO verify not updating duplicate email,
	// update in db
	// sql := "UPDATE users SET name = ?,email = ?,phone = ?,verified = ? WHERE uuid = ?;"
	// _, err = h.db.Exec(sql, req.Name, req.Email, req.Phone, req.Verified, req.UUID)

	if err != nil {
		h.sugaredLogger.Errorw("Error adding new user", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding new user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Successfully updated user",
		"req": req,
	})

}

func validateUpdateUserRequest(ctx *gin.Context) (UpdateUserRequest, error) {
	var req UpdateUserRequest

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
