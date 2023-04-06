package handler

import (
	"errors"
	"net/http"
	"time"

	maria_db "example.com/morethanjustlinks/db"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerService struct {
	maria_repo    maria_db.DbInterface
	sugaredLogger *zap.SugaredLogger
}

var secret = []byte("secret")
var PING_DB_ATTEMPTS = 90

func NewHandlerService(maria_repo maria_db.DbInterface, sugaredLogger *zap.SugaredLogger, attempts int) (*HandlerService, error) {

	for i := 0; i < attempts; i++ {
		if err := maria_repo.Ping(); err == nil {
			return &HandlerService{
				maria_repo:    maria_repo,
				sugaredLogger: sugaredLogger,
			}, nil
		}
		time.Sleep(time.Second)
	}

	return &HandlerService{}, errors.New("ping db timed out")

	// return &HandlerService{
	// 	maria_repo:    maria_repo,
	// 	sugaredLogger: sugaredLogger,
	// }, nil
}

func (h *HandlerService) SetupHandlerServiceRoutes() *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	cookieStore := sessions.NewCookieStore(secret)

	router.Use(
		sessions.Sessions("mysession", cookieStore),
	)

	// Serve frontend static files

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.LoadHTMLGlob("./views/templates/*")

	router.POST("login", h.Login)
	router.GET("logout", h.Logout)
	router.GET("newAccountForm", func(ctx *gin.Context) {
		// render html template for creating new user accounts
		ctx.HTML(http.StatusOK, "create_account.tmpl", gin.H{})
	})
	router.POST("newAccount", h.InsertUser)

	api := router.Group("/api")
	api.Use(h.Authentication)
	{
		api.GET("/", func(ctx *gin.Context) {
			session := sessions.Default(ctx)
			var count int
			v := session.Get("count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count += 1
			}
			session.Set("count", count)
			session.Save()
			ctx.JSON(http.StatusOK, gin.H{
				"message": count,
			})
		})
		api.GET("/:name/profile", func(ctx *gin.Context) {
			session := sessions.Default(ctx)
			var count int
			v := session.Get("count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count += 1
			}
			session.Set("count", count)
			session.Save()
			ctx.HTML(http.StatusOK, "profile.tmpl", gin.H{
				"count": count,
			})
		})
		api.GET("/:name/profile/edit", func(ctx *gin.Context) {
			session := sessions.Default(ctx)
			var count int
			v := session.Get("count")

			username := ctx.Param("name")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count += 1
			}
			session.Set("count", count)
			session.Save()
			ctx.HTML(http.StatusOK, "profile.tmpl", gin.H{
				"count": count,
				"name":  username,
			})
		})
		api.GET("/setup", h.SetupService)
		api.GET("/getAllUsers", h.GetAllUsers)

		api.POST("/delete", h.DeleteUser)
		api.POST("/update", h.UpdateUser)
	}

	return router
}
