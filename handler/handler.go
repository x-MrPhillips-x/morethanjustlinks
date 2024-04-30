package handler

import (
	"time"

	svcDB "example.com/morethanjustlinks/db"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	sugaredLogger *zap.SugaredLogger
	db            *gorm.DB
}

var secret = []byte("secret")
var PING_DB_ATTEMPTS = 90

// NewHandler holds all the dependencies for the service
func NewHandler(db *gorm.DB, sugaredLogger *zap.SugaredLogger, attempts int) (*Handler, error) {
	var err error
	var sqldb svcDB.SqlDB

	sqldb, err = db.DB()
	if err != nil {
		return &Handler{}, err
	}

	for i := 0; i < attempts; i++ {
		if err = sqldb.Ping(); err == nil {
			return &Handler{
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
	gin.SetMode(gin.ReleaseMode)

	cookieStore := sessions.NewCookieStore(secret)

	router.Use(
		sessions.Sessions("mysession", cookieStore),
	)

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	// router.GET("/setup", h.SetupService)
	// router.POST("login", h.Login)
	// router.GET("logout", h.Logout)
	router.POST("newAccount", h.NewAccount)
	router.GET("getAllUsers", h.GetAllUsers)
	router.POST("deleteUser", h.DeleteUser)
	router.GET("getUser", h.GetUserByName)

	// router.GET("/:name", h.GetProfile)

	// todo profile router use authentication
	// does pfp need to use sessions also?
	// pfp := router.Group("/:name/profile", h.Authentication)
	// pfp.POST("/edit", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"msg": "place holder for profile updates",
	// 	})
	// })

	// TODO and ^ behind pfp
	// router.POST("/deleteUser", h.DeleteUser)
	// router.POST("/deleteLink", h.DeleteLink)
	router.POST("/update", h.UpdateUser)

	// 		username := ctx.Param("name")
	// 		if v == nil {
	// 			count = 0
	// 		} else {
	// 			count = v.(int)
	// 			count += 1
	// 		}
	// 		session.Set("count", count)
	// 		session.Save()
	// 		ctx.HTML(http.StatusOK, "profile.tmpl", gin.H{
	// 			"count": count,
	// 			"name":  username,
	// 		})
	// 	})
	// 	api.GET("/setup", h.SetupService)
	// 	api.POST("/delete", h.DeleteUser)
	// 	api.POST("/update", h.UpdateUser)
	// }

	return router
}
