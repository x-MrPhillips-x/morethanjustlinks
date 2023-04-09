package handler

import (
	"fmt"
	"net/http"

	"example.com/morethanjustlinks/db"
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

type UserLink struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

func (h *HandlerService) GetProfile(ctx *gin.Context) {
	// get the name from the path
	name := ctx.Param("name")

	// get session data
	session := sessions.Default(ctx)
	uuid := session.Get("uuid")

	if name == "" || uuid == "" {
		h.sugaredLogger.Error("missing required profile params")
	}

	fmt.Println("This is the name:", name)
	fmt.Println("This is the session uuid:", uuid)

	queryStr := "select * from links where username = ?;"
	rows, err := h.maria_repo.Query(queryStr, name)
	if err != nil {
		h.sugaredLogger.Errorw("error getting profile data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
		return
	}

	var resp []UserLink
	for rows.Next() {
		var r UserLink
		if err := rows.Scan(&r.Username, &r.UUID, &r.Name, &r.Url); err != nil {
			h.sugaredLogger.Errorw("error adopting links to response", zap.Error(err))
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
			return
		}
		resp = append(resp, r)
	}

	if err := rows.Err(); err != nil {
		h.sugaredLogger.Errorw("error adopting rows", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})

		return
	}

	// get profile details
	ctx.JSON(http.StatusOK, gin.H{
		"links": resp,
	})
}

func (h *HandlerService) Login(ctx *gin.Context) {

	defer func() {
		h.sugaredLogger.Desugar().Sync() // flushes buffer, if any
	}()

	ctx.Header("Content-Type", "application/json")

	// Bind the input data
	var userAuth user.Auth
	var err error
	if err = ctx.BindJSON(&userAuth); err != nil {
		h.sugaredLogger.Errorw("error adapting request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if userAuth.Name == "" || userAuth.Psword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please enter required username and password"})
		return
	}

	var foundUser user.User
	if err := h.maria_repo.QueryRow(
		LOGIN_USER_QUERY, userAuth.Name).
		Scan(&foundUser.UUID, &foundUser.Name, &foundUser.Verified, &foundUser.Psword); err != nil {

		h.sugaredLogger.Errorw("account probably does not exists", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Check password hash + salt
	match := CheckPasswordHash(userAuth.Psword, foundUser.Psword)

	if !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	// set the session data
	session := sessions.Default(ctx)
	session.Set("uuid", uuid.New().String())
	session.Save()
	// if err := session.Save(); err != nil {
	// 	h.sugaredLogger.Errorw("error saving session data", zap.Error(err))
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{"msg": "successful login", "user": foundUser.Name, "verified": foundUser.Verified})
}

func adaptRowsToGetAllUsersResponse(rows db.RowsInterface) ([]GetAllUsersResponse, error) {
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "not an authorized user"})
		return
	}
	ctx.Next()
}

func (h *HandlerService) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"msg": "successful logout"})

}
