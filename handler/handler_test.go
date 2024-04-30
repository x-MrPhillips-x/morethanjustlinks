package handler

import (
	"database/sql"
	"testing"

	"example.com/morethanjustlinks/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type HandlerTestSuite struct {
	suite.Suite
	handler *Handler
	gormdb  *gorm.DB
	mock    sqlmock.Sqlmock
	router  *gin.Engine
}

func (h *HandlerTestSuite) SetupTest() {
	var err error
	var sqlDB *sql.DB

	sqlDB, h.mock, _ = sqlmock.New()

	dialector := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	// TODO mock migration?
	// expectedSQL := ".+"
	// h.mock.ExpectQuery(expectedSQL).WithArgs(1, 2).WillReturnError(nil)

	h.gormdb, err = db.NewGormDB(dialector, &gorm.Config{})

	assert.Nil(h.T(), err)

	h.handler, err = NewHandler(h.gormdb, zap.NewNop().Sugar(), 3)
	assert.Nil(h.T(), err)

	h.router = h.handler.SetupHandlerRoutes()
	assert.NotNil(h.T(), h.router)

}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
