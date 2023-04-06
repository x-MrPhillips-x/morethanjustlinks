package handler

import (
	"fmt"
	"net/http"

	"example.com/morethanjustlinks/user"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	SELECT_ALL_USERS = "SELECT uuid,name,email,phone,verified FROM users;"
)

type GetAllUsersResponse struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Verified string `json:"verified"`
}

type ReadInterface interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
	Close() error
}

func (h *HandlerService) GetAllUsers(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	rows, err := h.maria_repo.Query(SELECT_ALL_USERS)
	if err != nil {
		h.sugaredLogger.Errorw("Error fetching all users", zap.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching all users"})
		return
	}

	defer func() {
		h.sugaredLogger.Desugar().Sync()
		rows.Close()
	}()

	resp, err := adaptRowsToGetAllUsersResponse(rows)
	if err != nil {
		h.sugaredLogger.Errorw("Error adapting all users to response", zap.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error adapting all user to response"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Successfully fetched all users",
		"resp": resp,
	})

}

func (h HandlerService) Login(ctx *gin.Context) {

	defer func() {
		h.sugaredLogger.Desugar().Sync() // flushes buffer, if any
	}()

	session := sessions.Default(ctx)
	ctx.Header("Content-Type", "application/json")

	// Bind the input data
	var userAuth user.Auth
	if err := ctx.BindJSON((&userAuth)); err != nil {
		h.sugaredLogger.Errorw("Error binding fronted json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong with your login request..."})
		return
	}

	msg := fmt.Sprintf("given username %v and pass %v", userAuth.Username, userAuth.Psword)

	// set the session data
	session.Set(user.USERKEY, userAuth.Username)
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": msg})
}

func adaptRowsToGetAllUsersResponse(rows ReadInterface) ([]GetAllUsersResponse, error) {
	var resp []GetAllUsersResponse
	for rows.Next() {
		var r GetAllUsersResponse
		if err := rows.Scan(&r.UUID, &r.Name, &r.Email, &r.Phone, &r.Verified); err != nil {
			return resp, err
		}
		resp = append(resp, r)
	}

	if err := rows.Err(); err != nil {
		return resp, err
	}

	return resp, nil
}
