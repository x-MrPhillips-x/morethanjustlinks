package handler

import (
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

	req, err := validateUpdateUserRequest(ctx)
	if err != nil {
		h.sugaredLogger.Errorw("invalid update user request",
			zap.Any("error", err.Error()),
			zap.Any("req", req))
		ctx.JSON(http.StatusBadRequest, handleValidationErrors((err)))
		return
	}

	user := &models.User{ID: req.ID}

	result := h.db.First(user)
	if result.Error != nil {
		h.sugaredLogger.Errorw("error fetching user to update", zap.Any("error", result.Error.Error()), zap.Any("user", user))
		msg := "something went wrong..."
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": msg})
		return
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Phone = req.Phone
	user.Role = req.Role

	result = h.db.Save(user)

	if result.Error != nil {
		h.sugaredLogger.Errorw("error updating user", zap.Any("error", result.Error.Error()), zap.Any("user", user))
		msg := "something went wrong..."
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}
