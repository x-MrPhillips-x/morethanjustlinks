package handler

import (
	"time"

	"example.com/morethanjustlinks/config"
	svcDB "example.com/morethanjustlinks/db"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	appConfig     config.AppConfig
	sugaredLogger *zap.SugaredLogger
	db            *gorm.DB
}

var PING_DB_ATTEMPTS = 90

// NewHandler holds all the dependencies for the service
func NewHandler(appConfig config.AppConfig, db *gorm.DB, sugaredLogger *zap.SugaredLogger, attempts int) (*Handler, error) {
	var err error
	var sqldb svcDB.SqlDB

	sqldb, err = db.DB()
	if err != nil {
		return &Handler{}, err
	}

	for i := 0; i < attempts; i++ {
		if err = sqldb.Ping(); err == nil {
			return &Handler{
				appConfig:     appConfig,
				sugaredLogger: sugaredLogger,
				db:            db,
			}, nil
		}
		time.Sleep(time.Second)
	}

	return &Handler{}, err
}

func (h *Handler) SetupHandlerRoutes() *gin.Engine {

	router := gin.Default()

	// for file up loads
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	gin.SetMode(gin.ReleaseMode)

	cookieStore := sessions.NewCookieStore(h.appConfig.Server.Sessions)

	// router.Use(cors.New(config))
	router.Use(
		sessions.Sessions("mysession", cookieStore),
	)

	router.POST("login", h.Login)
	router.POST("logout", h.Logout)
	router.POST("newAccount", h.NewAccount)
	router.GET("getUser", h.GetUserByID)
	router.GET("getAllUsers", h.AuthMiddleware(), h.GetAllUsers)
	router.POST("deleteUser", h.AuthMiddleware(), h.DeleteUser)
	router.POST("upload", h.AuthMiddleware(), h.Upload)
	router.POST("update", h.AuthMiddleware(), h.UpdateUser)
	router.GET("incr", func(c *gin.Context) {
		session := sessions.Default(c)
		c.Header("Content-Type", "application/json")

		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	return router
}
