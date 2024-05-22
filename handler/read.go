package handler

import (
	"net/http"

	"example.com/morethanjustlinks/models"
	"example.com/morethanjustlinks/user"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) GetUserByID(ctx *gin.Context) {
	var req models.GetUserRequest
	var user models.User

	req.ID = ctx.Query("id")

	result := h.db.First(&user, "id = ?", req.ID)

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": user})
}

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	var users []models.User
	session := sessions.Default(ctx)
	count := session.Get("count")
	result := h.db.Find(&users)
	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": users, "count": count})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": []models.User{}})

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

func (h *Handler) Login(ctx *gin.Context) {

	defer func() {
		h.sugaredLogger.Desugar().Sync() // flushes buffer, if any
	}()

	ctx.Header("Content-Type", "application/json")

	// Bind the input data
	var req user.Auth
	var err error
	if err = ctx.BindJSON(&req); err != nil {
		h.sugaredLogger.Errorw("error adapting request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong..."})
		return
	}

	if req.Email == "" || req.Psword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please enter required username and password"})
		return
	}

	var foundUser models.User
	h.db.Where("email = ? ", req.Email).First(&foundUser)

	// Check password hash + salt
	if match := CheckPasswordHash(req.Psword, foundUser.Psword); !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	// set the session data
	// for nextjs middleware
	session := sessions.Default(ctx)
	var uuid int

	if v := session.Get("uuid"); v == nil {
		uuid = 1
		session.Set("uuid", uuid)
		session.Save()
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successful login", "user": foundUser.Name, "verified": foundUser.Verified, "uuid": uuid})
}

func (h *Handler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	count0 := session.Get("count")
	session.Clear()
	session.Save()
	count1 := session.Get("count")

	// TODO update DB

	ctx.JSON(http.StatusOK, gin.H{"count0": count0, "count1": count1, "message": "successful logout"})

}
