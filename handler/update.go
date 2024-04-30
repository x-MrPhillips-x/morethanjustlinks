package handler

import (
	"fmt"
	"net/http"

	"example.com/morethanjustlinks/models"
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

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var req models.User
	var user models.User

	if err := ctx.BindJSON(&req); err != nil {
		h.sugaredLogger.Errorw("invalid update user request",
			zap.Any("error", err.Error()),
			zap.Any("req", req))

		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if !isValidEmail(req.Email) {
		msg := fmt.Sprintf("%v is not a valid email address", req.Email)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		return
	}

	if !isValidPhoneNumber(req.Phone) {
		msg := fmt.Sprintf("%v is not a valid mobile number", req.Email)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		return
	}

	if !isValidUsername(req.Name) {
		msg := fmt.Sprintf("%v is not a valid mobile number", req.Email)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		return
	}

	result := h.db.Model(&user).Updates(models.User{
		Phone:    req.Phone,
		Email:    req.Email,
		Verified: req.Verified,
		Role:     req.Role,
	})

	if result.Error != nil {
		h.sugaredLogger.Errorw("error updating user", zap.Any("error", result.Error.Error()), zap.Any("user", user))
		msg := "something went wrong..."
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": msg})
	}
}
