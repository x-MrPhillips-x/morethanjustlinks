package db

import (
	"database/sql"
	"errors"
	"testing"

	"example.com/morethanjustlinks/mocks"
	"github.com/DATA-DOG/go-sqlmock"
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

func TestQueryReturnedRows(t *testing.T) {
	db, sqlmock, _ := sqlmock.New()

	sqlmock.ExpectQuery("queryStr").WillReturnRows(sqlmock.NewRows([]string{"uuid"}).AddRow("some-uuid"))

	rows, err := Query(db, "queryStr")

	assert.Nil(t, err)

	assert.NotNil(t, rows)

	sqlmock.ExpectationsWereMet()

}

func TestQueryRowsNextTrue(t *testing.T) {
	db, sqlmock, _ := sqlmock.New()

	sqlmock.ExpectQuery("queryStr").WillReturnRows(sqlmock.NewRows([]string{"uuid"}).AddRow("some-uuid"))

	rows_under_test, err := Query(db, "queryStr")

	assert.True(t, Next(rows_under_test))

	assert.Nil(t, err)

	assert.NotNil(t, rows_under_test)

	sqlmock.ExpectationsWereMet()

}

func TestQueryRowsNextFalse(t *testing.T) {
	db, sqlmock, _ := sqlmock.New()

	sqlmock.ExpectQuery("queryStr").WillReturnRows(sqlmock.NewRows([]string{}))

	rows_under_test, err := Query(db, "queryStr")

	assert.False(t, Next(rows_under_test))

	assert.Nil(t, err)

	assert.NotNil(t, rows_under_test)

	sqlmock.ExpectationsWereMet()

}

type SomeTestStruct struct {
	UUID string `json:"uuid"`
}

func TestQueryRowsScanUnsupportedScanError(t *testing.T) {
	var resp SomeTestStruct
	db, sqlmock, _ := sqlmock.New()

	sqlmock.ExpectQuery("queryStr").WillReturnRows(sqlmock.NewRows([]string{"hair"}).AddRow("some-wig"))

	rows_under_test, _ := Query(db, "queryStr")

	assert.True(t, Next(rows_under_test))

	assert.NotNil(t, Scan(rows_under_test, &resp))

	sqlmock.ExpectationsWereMet()

}

// Todo the row is not texted correctly
