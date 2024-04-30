package handler

import (
	"fmt"
	"net/http"

	"example.com/morethanjustlinks/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) GetUserByName(ctx *gin.Context) {
	var req models.GetUserRequest
	var user models.User

	if err := ctx.BindJSON(&req); err != nil {
		h.sugaredLogger.Errorw("error getting user",
			zap.Any("error", err.Error()),
			zap.Any("request", req))

		msg := fmt.Sprintf("%v is not a valid user", req.Name)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
	}

	result := h.db.First(&user, "name = ?", req.Name)

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": user})
}

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var users []models.User
	result := h.db.Find(&users)
	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusOK, users)
		return
	}
	ctx.JSON(http.StatusOK, []models.User{})
}

type UserLink struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

// func (h *HandlerService) GetProfile(ctx *gin.Context) {
// 	// get the name from the path
// 	name := ctx.Param("name")

// 	// get session data
// 	session := sessions.Default(ctx)
// 	uuid := session.Get("uuid")

// 	if name == "" || uuid == "" {
// 		h.sugaredLogger.Error("missing required profile params")
// 	}

// 	queryStr := "select * from links where username = ?;"
// 	rows, err := h.maria_repo.Query(queryStr, name)

// 	if err != nil {
// 		h.sugaredLogger.Errorw("error getting profile data", zap.Error(err))
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
// 		return
// 	}

// 	var resp []UserLink
// 	for rows.Next() {
// 		var r UserLink
// 		if err := rows.Scan(&r.Username, &r.UUID, &r.Name, &r.Url); err != nil {
// 			h.sugaredLogger.Errorw("error adopting links to response", zap.Error(err))
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})
// 			return
// 		}
// 		resp = append(resp, r)
// 	}

// 	if err := rows.Err(); err != nil {
// 		h.sugaredLogger.Errorw("error adopting rows", zap.Error(err))
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong..."})

// 		return
// 	}

// 	// get profile details
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"links": resp,
// 	})
// }

// func (h *HandlerService) Login(ctx *gin.Context) {

// 	defer func() {
// 		h.sugaredLogger.Desugar().Sync() // flushes buffer, if any
// 	}()

// 	ctx.Header("Content-Type", "application/json")

// 	// Bind the input data
// 	var userAuth user.Auth
// 	var err error
// 	if err = ctx.BindJSON(&userAuth); err != nil {
// 		h.sugaredLogger.Errorw("error adapting request", zap.Error(err))
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong..."})
// 		return
// 	}

// 	if userAuth.Name == "" || userAuth.Psword == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please enter required username and password"})
// 		return
// 	}

// 	var foundUser user.User
// 	if err := h.maria_repo.QueryRow(
// 		LOGIN_USER_QUERY, userAuth.Name).
// 		Scan(&foundUser.UUID, &foundUser.Name, &foundUser.Verified, &foundUser.Psword); err != nil {

// 		h.sugaredLogger.Errorw("account probably does not exists", zap.Error(err))
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}

// 	// Check password hash + salt
// 	match := CheckPasswordHash(userAuth.Psword, foundUser.Psword)

// 	if !match {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
// 		return
// 	}

// 	// set the session data
// 	session := sessions.Default(ctx)
// 	session.Set("uuid", uuid.New().String())
// 	session.Save()
// 	// if err := session.Save(); err != nil {
// 	// 	h.sugaredLogger.Errorw("error saving session data", zap.Error(err))
// 	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 	// 	return
// 	// }

// 	ctx.JSON(http.StatusOK, gin.H{"msg": "successful login", "user": foundUser.Name, "verified": foundUser.Verified})
// }

// func adaptRowsToGetAllUsersResponse(rows db.RowsInterface) ([]GetAllUsersResponse, error) {
// 	var resp []GetAllUsersResponse
// 	for rows.Next() {
// 		var r GetAllUsersResponse
// 		if err := rows.Scan(&r.UUID, &r.Name, &r.Email, &r.Phone, &r.Verified); err != nil {
// 			return resp, err
// 		}
// 		resp = append(resp, r)
// 	}

// 	return resp, nil
// }

// func (h *HandlerService) Authentication(ctx *gin.Context) {
// 	session := sessions.Default(ctx)
// 	sessionUUID := session.Get("uuid")

// 	if sessionUUID == nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "not an authorized user"})
// 		return
// 	}
// 	ctx.Next()
// }

// func (h *HandlerService) Logout(ctx *gin.Context) {
// 	session := sessions.Default(ctx)
// 	session.Clear()
// 	session.Save()

// 	ctx.JSON(http.StatusOK, gin.H{"msg": "successful logout"})

// }
