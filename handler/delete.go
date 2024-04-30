package handler

import (
	"fmt"
	"net/http"

	"example.com/morethanjustlinks/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DeleteUserRequest struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type DeleteLinkRequest struct {
	UUID string `json:"uuid"`
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		h.sugaredLogger.Errorw("Error not a valid delete user request",
			zap.Any("error", err.Error()),
			zap.Any("user", user))
		msg := fmt.Sprintf("%s could not be deleted : %s", user.Name, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		return
	}

	result := h.db.Where("ID = ?", user.ID).Delete(&user)
	if result.Error != nil {
		h.sugaredLogger.Errorw("failed to delete user",
			zap.Any("error", result.Error.Error()),
			zap.Any("user", user))
		msg := fmt.Sprintf("%s failed to delete", user.Name)
		ctx.JSON(http.StatusOK, gin.H{"msg": msg})
		return
	}

	if result.RowsAffected == 0 {
		msg := fmt.Sprintf("%s failed to delete", user.Name)
		ctx.JSON(http.StatusOK, gin.H{"msg": msg})
		return
	}

	msg := fmt.Sprintf("%s was successfully deleted", user.Name)
	ctx.JSON(http.StatusOK, gin.H{"msg": msg})
}
