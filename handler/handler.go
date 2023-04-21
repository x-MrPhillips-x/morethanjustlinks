package handler

import (
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
	var err error
	for i := 0; i < attempts; i++ {
		if err = maria_repo.Ping(); err == nil {
			return &HandlerService{
				maria_repo:    maria_repo,
				sugaredLogger: sugaredLogger,
			}, nil
		}
		time.Sleep(time.Second)
	}

	return &HandlerService{}, err
}

func (h *HandlerService) SetupHandlerServiceRoutes() *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	cookieStore := sessions.NewCookieStore(secret)

	router.Use(
		sessions.Sessions("mysession", cookieStore),
	)

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend", true)))
	router.LoadHTMLGlob("./frontend/components/*")

	router.GET("/setup", h.SetupService)
	router.POST("/login", h.Login)
	router.GET("/logout", h.Logout)
	router.GET("/newAccountForm", h.GetNewAccountForm)
	router.POST("/newAccount", h.NewAccount)
	router.GET("/:name", h.GetProfile)
	router.GET("/getAllUsers", h.GetAllUsers)

	// todo profile router use authentication
	// does pfp need to use sessions also?
	pfp := router.Group("/pfp")
	// pfp.Use(h.Authentication)
	pfp.GET("/edit", h.EditProfileForm)
	pfp.POST("/edit", h.CreateLink)

	// TODO and ^ behind pfp
	pfp.POST("/deleteUser", h.DeleteUser)
	pfp.POST("/deleteLink", h.DeleteLink)
	pfp.POST("/update", h.UpdateUser)

	return router
}
