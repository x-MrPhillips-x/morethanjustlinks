package db

import (
	"testing"

	"example.com/morethanjustlinks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DBInterfaceImplementationsSuite struct {
	suite.Suite
	db      *mocks.DbInterface
	db_rows *mocks.RowsInterface
}

func (suite *DBInterfaceImplementationsSuite) SetupTest() {
	suite.db = &mocks.DbInterface{}
	suite.db_rows = &mocks.RowsInterface{}
}

func TestDBInterfaceImplementationsSuite(t *testing.T) {
	suite.Run(t, new(DBInterfaceImplementationsSuite))
}

func TestPingSuccess(t *testing.T) {
	db := &mocks.DbInterface{}
	db.On("Ping").Return(nil)

	err := Ping(db)

	assert.Nil(t, err)

	db.AssertExpectations(t)
}

type SomeTestStruct struct {
	UUID string `json:"uuid"`
}
