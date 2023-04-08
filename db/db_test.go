package db

import (
	"database/sql"
	"errors"
	"testing"

	"example.com/morethanjustlinks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestConnect(t *testing.T) {
	db, err := Connect()
	assert.Nil(t, err)
	assert.NotNil(t, db)

}

func TestPingSuccess(t *testing.T) {
	db := &mocks.DbInterface{}
	db.On("Ping").Return(nil)

	err := Ping(db)

	assert.Nil(t, err)

	db.AssertExpectations(t)
}

func TestExec(t *testing.T) {
	db := &mocks.DbInterface{}

	db.On("Exec", "query", mock.Anything).Return(nil, nil)

	_, err := Exec(db, "query")

	assert.Nil(t, err)

	db.AssertExpectations(t)

}

func TestExecError(t *testing.T) {
	db := &mocks.DbInterface{}

	db.On("Exec", "query", mock.Anything).Return(nil, errors.New("some Error"))

	_, err := Exec(db, "query")

	assert.NotNil(t, err)
	assert.Equal(t, errors.New("some Error"), err)

	db.AssertExpectations(t)

}

func TestQuery(t *testing.T) {
	db := &mocks.DbInterface{}

	db.On("Query", "querystr", mock.Anything).Return(&sql.Rows{}, nil)

	rows, err := Query(db, "querystr")

	assert.Nil(t, err)
	assert.Equal(t, &sql.Rows{}, rows)

	db.AssertExpectations(t)

}

func TestQueryError(t *testing.T) {
	db := &mocks.DbInterface{}

	db.On("Query", "querystr", mock.Anything).Return(nil, errors.New("some Error"))

	rows, err := Query(db, "querystr")

	assert.Nil(t, rows)
	assert.Equal(t, errors.New("some Error"), err)

	db.AssertExpectations(t)

}
