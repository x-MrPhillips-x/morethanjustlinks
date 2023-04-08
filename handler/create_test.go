package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

			if tt.expectCode == 200 {
				h.db_mock.On(
					"Exec",
					"DROP TABLE IF EXISTS users;").Return(sqlmock.NewResult(1, 1), nil).Once()
				h.db_mock.On(
					"Exec",
					CREATE_USERS_TABLE).Return(sqlmock.NewResult(1, 1), nil).Once()

			}

			if tt.isDeleteErr {
				h.db_mock.On(
					"Exec",
					"DROP TABLE IF EXISTS users;").Return(sqlmock.NewResult(1, 1), errors.New("some error")).Once()
			}

			if tt.isCreateErr {
				h.db_mock.On(
					"Exec",
					"DROP TABLE IF EXISTS users;").Return(sqlmock.NewResult(1, 1), nil).Once()
				h.db_mock.On(
					"Exec",
					CREATE_USERS_TABLE).Return(sqlmock.NewResult(1, 1), errors.New("some error")).Once()
			}

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
	}{
		// {
		// 	"Error with input data",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"please enter the required request fields",
		// 	400,
		// },
		// {
		// 	"Error with input data - username is not valid",
		// 	[]byte(`{
		// 		"name": "m",
		// 		"email": "morpheus@mail.com",
		// 		"phone":"7777777777",
		// 		"psword":"leader"
		// 	}`),
		// 	"error",
		// 	"please enter the required request fields",
		// 	400,
		// },
		// {
		// 	"Error with input data - email format",
		// 	[]byte(`{
		// 		"name": "morpheus",
		// 		"email": "morpheus",
		// 		"phone":"7777777777",
		// 		"psword":"leader"
		// 	}`),
		// 	"error",
		// 	"please enter the required request fields",
		// 	400,
		// },
		// {
		// 	"Error with input data - phone # min digits",
		// 	[]byte(`{
		// 		"name": "morpheus",
		// 		"email": "morpheus@mail.com",
		// 		"phone":"77777",
		// 		"psword":"leader"
		// 	}`),
		// 	"error",
		// 	"please enter the required request fields",
		// 	400,
		// },
		// {
		// 	"Error with input data - phone # max digits",
		// 	[]byte(`{
		// 		"name": "morpheus",
		// 		"email": "morpheus@mail.com",
		// 		"phone":"777777777777777",
		// 		"psword":"leader"
		// 	}`),
		// 	"error",
		// 	"please enter the required request fields",
		// 	400,
		// },
		{
			"Success creating a new account",
			[]byte(`{
				"name": "morpheus",
				"email": "morpheus@mail.com",
				"phone":"7777777777",
				"psword":"leader"
			}`),
			"msg",
			"successfully created new user",
			200,
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
		// 	"msg",
		// 	"successfully created new user",
		// 	200,
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

			if tt.expectCode == 200 {
				// var rows *sql.Rows
				h.rows_mock.On("Next").Return(false)
				h.rows_mock.On(
					"Scan",
					mock.Anything,
				).Return(false, nil)
				h.db_mock.On(
					"Query",
					"SELECT COUNT(*) FROM users WHERE name = ?",
					"morpheus",
					mock.Anything).Return(h.rows_mock, nil)
				h.db_mock.On(
					"Exec",
					"INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)",
					mock.Anything,
					"morpheus",
					"morpheus@mail.com",
					"7777777777",
					"leader",
					mock.Anything).Return(sqlmock.NewResult(1, 1), nil).Once()
			}

			// if tt.expectCode == 400 {
			// 	var rows *sql.Rows
			// 	h.db_mock.On(
			// 		"Query",
			// 		"SELECT COUNT(*) FROM users WHERE name = ?",
			// 		"morpheus",
			// 		mock.Anything).Return(rows, errors.New("some DB")).Once()
			// }

			if tt.expectCode == 500 {
				h.db_mock.On(
					"Exec",
					"INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)",
					mock.Anything,
					"morpheus",
					"morpheus@mail.com",
					"7777777777",
					"leader",
					mock.Anything).Return(sqlmock.NewResult(0, 0), errors.New("some DB")).Once()
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/newAccount", bytes.NewBuffer(tt.reqParams))

			h.router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			fmt.Println("What is the response?", w.Body.String())

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
