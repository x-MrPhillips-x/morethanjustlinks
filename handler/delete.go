package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteUserRequest struct {
	Name string `json:"name"`
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
