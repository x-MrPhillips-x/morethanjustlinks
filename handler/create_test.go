package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
		mocks        func()
	}{
		// {
		// 	"Error with input data",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"Error with input data",
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
		// 	"Error with input data",
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
		// 	"Error with input data",
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
		// 	"Error with input data",
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
		// 	"Error with input data",
		// 	400,
		// },
		// {
		// 	"Success creating a new account",
		// 	[]byte(`{
		// 		"name": "morpheus",
		// 		"email": "morpheus@mail.com",
		// 		"phone":"7777777777",
		// 		"psword":"leader"
		// 	}`),
		// 	"msg",
		// 	"successfully created new user",
		// 	200,
		// 	func() {
		// 		h.db_mock.On(
		// 			"Exec",
		// 			"INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)",
		// 			mock.Anything,
		// 			"morpheus",
		// 			"morpheus@mail.com",
		// 			"7777777777",
		// 			mock.Anything,
		// 			mock.Anything).Return(sqlmock.NewResult(1, 1), nil).Once()
		// 	},
		// },
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
		// 	func() {
		// 		h.db_mock.On(
		// 			"Exec",
		// 			"INSERT INTO users (uuid,name,email,phone,psword,verified) VALUES (?,?,?,?,?,?)",
		// 			mock.Anything,
		// 			"morpheus",
		// 			"morpheus@mail.com",
		// 			"7777777777",
		// 			mock.Anything,
		// 			mock.Anything).Return(sqlmock.NewResult(1, 1), nil).Once()
		// 	},
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
