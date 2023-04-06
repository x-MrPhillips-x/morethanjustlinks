package handler

import (
	"errors"
	"net/http"
	"time"

	maria_db "example.com/morethanjustlinks/db"
	"github.com/gin-gonic/contrib/sessions"
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

	router.Use(
		sessions.Sessions("mysession", sessions.NewCookieStore(secret)),
	)

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
	})

	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "on me",
		})
	})
	router.GET("/setup", h.SetupService)
	router.GET("/getAllUsers", h.GetAllUsers)

	router.POST("/login", h.Login)
	router.POST("/newAccount", h.InsertUser)
	router.POST("/delete", h.DeleteUser)
	router.POST("/update", h.UpdateUser)

	return router
}
