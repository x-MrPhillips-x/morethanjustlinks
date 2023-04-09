package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteUserRequest struct {
	Name string `json:"name"`
}

type DeleteLinkRequest struct {
	UUID string `json:"uuid"`
}

func (h *HandlerService) DropUsersTable(ctx *gin.Context) {
	_, err := h.maria_repo.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error dropping users"})
		ctx.Error(err)
	}
	ctx.Next()
}

func (h *HandlerService) DropLinksTable(ctx *gin.Context) {
	_, err := h.maria_repo.Exec("DROP TABLE IF EXISTS links;")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error dropping links"})
		ctx.Error(err)
	}
	ctx.Next()
}

func (h *HandlerService) DeleteUser(ctx *gin.Context) {
	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	var req DeleteUserRequest
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON((&req)); err != nil {
		h.sugaredLogger.Errorw("Error not a valid delete user request", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error not a valid delete user request"})
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error missing required name"})
		return
	}

	// delete from db
	sql := "DELETE FROM users WHERE name = ?"
	_, err := h.maria_repo.Exec(sql, req.Name)
	if err != nil {
		h.sugaredLogger.Errorw("Error deleting user", zap.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted user"})

}

func (h *HandlerService) DeleteLink(ctx *gin.Context) {
	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	var req DeleteLinkRequest
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON((&req)); err != nil {
		h.sugaredLogger.Errorw("Error not a valid delete link request", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error not a valid delete link request"})
		return
	}

	if req.UUID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error missing required uuid"})
		return
	}
	// delete from db
	sql := "DELETE FROM links WHERE uuid = ?"
	_, err := h.maria_repo.Exec(sql, req.UUID)
	if err != nil {
		h.sugaredLogger.Errorw("Error deleting link", zap.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove link"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted link"})

}
