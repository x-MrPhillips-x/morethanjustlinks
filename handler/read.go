package handler

import (
	"fmt"
	"net/http"

	"example.com/morethanjustlinks/user"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	SELECT_ALL_USERS = "SELECT uuid,name,email,phone,verified FROM users;"
	LOGIN_USER_QUERY = "SELECT uuid,name,verified,psword FROM users WHERE name= ?;"
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
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	defer func() {
		h.sugaredLogger.Desugar().Sync()
	}()

	fmt.Println("these are the rows:", rows)
	fmt.Println("these are the rows:", rows)

	var resp []GetAllUsersResponse
	for rows.Next() {
		var r GetAllUsersResponse
		if err := rows.Scan(&r.UUID, &r.Name, &r.Email, &r.Phone, &r.Verified); err != nil {
			h.sugaredLogger.Errorw("error adapting all users to response", zap.Any("error", err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
			break
		}
		resp = append(resp, r)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"resp": resp,
	})

}

func (h *HandlerService) Login(ctx *gin.Context) {

	defer func() {
		h.sugaredLogger.Desugar().Sync() // flushes buffer, if any
	}()

	session := sessions.Default(ctx)
	ctx.Header("Content-Type", "application/json")

	// Bind the input data
	var userAuth user.Auth
	if err := ctx.BindJSON(&userAuth); err != nil {
		h.sugaredLogger.Errorw("Error binding fronted json", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong with your login request..."})
		return
	}

	if userAuth.Name == "" || userAuth.Psword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please enter required username and password"})
		return
	}

	msg := fmt.Sprintf("given username %v and pass %v", userAuth.Name, userAuth.Psword)

	var foundUser user.User
	if err := h.maria_repo.QueryRow(
		LOGIN_USER_QUERY, userAuth.Name).
		Scan(&foundUser.UUID, &foundUser.Name, &foundUser.Verified, &foundUser.Psword); err != nil {

		h.sugaredLogger.Errorw("Error logging in", zap.Any("error", err), zap.String("user", userAuth.Name))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong with your login request...02"})
		return
	}

	// Check password hash + salt
	match := CheckPasswordHash(userAuth.Psword, foundUser.Psword)

	if !match {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized user"})
		return
	}

	// set the session data
	session.Set(user.UUIDKEY, uuid.New().String())
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to set session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": msg, "user": foundUser.Name, "verified": foundUser.Verified})
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

	// if err := rows.Err(); err != nil {
	// 	return resp, err
	// }

	return resp, nil
}

func (h *HandlerService) Authentication(ctx *gin.Context) {
	session := sessions.Default(ctx)
	sessionUUID := session.Get("uuid")

	if sessionUUID == nil {
		ctx.File("index.html")
		return
	}
}

func (h *HandlerService) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"msg": "successful logout"})

}
