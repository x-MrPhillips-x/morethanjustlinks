package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/morethanjustlinks/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (h *HandlerTestSuite) TestSetupService() {
	tests := []struct {
		name         string
		reqParams    []byte
		expectMsgKey string
		expectMsg    string
		expectCode   int
		isDeleteErr  bool
		isCreateErr  bool
	}{
		// {
		// 	"Error dropping users table",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"error dropping users",
		// 	500,
		// 	true,
		// 	false,
		// },
		// {
		// 	"Success dropping user table, but failed to create users",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"error creating users",
		// 	500,
		// 	false,
		// 	true,
		// },
		// {
		// 	"Happy path, users table dropped and create users table",
		// 	[]byte(`{}`),
		// 	"msg",
		// 	"created user tables succesfully",
		// 	200,
		// 	false,
		// 	false,
		// },
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			// if tt.expectCode == 200 {
			// 	h.db_mock.On(
			// 		"Exec",
			// 		"DROP TABLE IF EXISTS users;").Return(sqlmock.NewResult(1, 1), nil).Once()
			// 	h.db_mock.On(
			// 		"Exec",
			// 		CREATE_USERS_TABLE).Return(sqlmock.NewResult(1, 1), nil).Once()

			// }

			// if tt.isDeleteErr {
			// 	h.db_mock.On(
			// 		"Exec",
			// 		"DROP TABLE IF EXISTS users;").Return(sqlmock.NewResult(1, 1), errors.New("some error")).Once()
			// }

			// if tt.isCreateErr {
			// 	h.db_mock.On(
			// 		"Exec",
			// 		"DROP TABLE IF EXISTS users;").Return(sqlmock.NewResult(1, 1), nil).Once()
			// 	h.db_mock.On(
			// 		"Exec",
			// 		CREATE_USERS_TABLE).Return(sqlmock.NewResult(1, 1), errors.New("some error")).Once()
			// }

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", "/setup", bytes.NewBuffer(tt.reqParams))

			h.router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			var actualResponse map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			if tt.expectMsgKey == "error" {
				assert.Equal(t, tt.expectMsg, actualResponse["error"])
			} else {
				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
			}
		})
	}
}

func (h *HandlerTestSuite) TestNewAccountRoute() {
	tests := []struct {
		name         string
		reqParams    []byte
		expectMsgKey string
		expectMsg    string
		expectCode   int
		sqlmocks     sqlmock.Sqlmock
		hasRows      bool
	}{
		{
			"Error with input data",
			[]byte(`{}`),
			"error",
			"please enter the required request fields",
			400,
			h.mock,
			false,
		},
		{
			"Error with input data - username is not valid",
			[]byte(`{
				"name": "m",
				"email": "morpheus@mail.com",
				"phone":"7777777777",
				"psword":"leader"
			}`),
			"error",
			"please enter the required request fields",
			400,
			h.mock,
			false,
		},
		{
			"Error with input data - email format",
			[]byte(`{
				"name": "morpheus",
				"email": "morpheus",
				"phone":"7777777777",
				"psword":"leader"
			}`),
			"error",
			"please enter the required request fields",
			400,
			h.mock,
			false,
		},
		{
			"Error with input data - phone # min digits",
			[]byte(`{
				"name": "morpheus",
				"email": "morpheus@mail.com",
				"phone":"77777",
				"psword":"leader"
			}`),
			"error",
			"please enter the required request fields",
			400,
			h.mock,
			false,
		},
		{
			"Error with input data - phone # max digits",
			[]byte(`{
				"name": "morpheus",
				"email": "morpheus@mail.com",
				"phone":"777777777777777",
				"psword":"leader"
			}`),
			"error",
			"please enter the required request fields",
			400,
			h.mock,
			false,
		},
		{
			"Bad request user already exists",
			[]byte(`{
				"name": "morpheus",
				"email": "morpheus@mail.com",
				"phone":"7777777777",
				"psword":"leader"
			}`),
			"error",
			"user already exists",
			400,
			h.mock,
			true,
		},
		{
			"Something went wrong hashing and salting",
			[]byte(`{
				"name": "morpheus",
				"email": "morpheus@mail.com",
				"phone":"7777777777",
				"psword":"leader",
				"verified":true
			}`),
			"error",
			"something went wrong...",
			500,
			h.mock,
			true,
		},
		// {
		// 	"Success creating a new account - verified true",
		// 	[]byte(`{
		// 		"name": "morpheus",
		// 		"email": "morpheus@mail.com",
		// 		"phone":"7777777777",
		// 		"psword":"leader",
		// 		"verified":true
		// 	}`),
		// 	"error",
		// 	"successfully created new user",
		// 	500,
		// 	h.mock,
		// 	true,
		// },
		// {
		// 	"Error inserting user into DB",
		// 	[]byte(`{
		// 		"name": "morpheus",
		// 		"email": "morpheus@mail.com",
		// 		"phone":"7777777777",
		// 		"psword":"leader",
		// 		"verified":true
		// 	}`),
		// 	"error",
		// 	"Error adding new user",
		// 	500,
		// },
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/newAccount", bytes.NewBuffer(tt.reqParams))

			var mock_rows = sqlmock.NewRows([]string{})
			if tt.hasRows && tt.expectCode == 400 {
				mock_rows := sqlmock.NewRows([]string{"count"}).AddRow("1").AddRow("2")
				tt.sqlmocks.ExpectQuery("SELECT COUNT(.*) FROM users WHERE name = ?").WillReturnRows(mock_rows)
			}

			if tt.hasRows {
				var mock_hasher mocks.HasherInterface
				tt.sqlmocks.ExpectQuery("SELECT COUNT(.*) FROM users WHERE name = ?").WillReturnRows(mock_rows)
				mock_hasher.On("HashPassword", "leader").Return("leader+hash+salt", nil)
				tt.sqlmocks.ExpectExec("INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)").WillReturnResult(sqlmock.NewErrorResult(errors.New("some error")))
			}

			h.router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			var actualResponse map[string]string
			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			if tt.expectMsgKey == "error" {
				assert.Equal(t, tt.expectMsg, actualResponse["error"])
			} else {
				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
			}
		})
	}
}
