package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (h *HandlerTestSuite) TestGetAllUsers() {
	tests := []struct {
		name     string
		mock     sqlmock.Sqlmock
		queryStr string
		rows     *sqlmock.Rows
		queryErr error
	}{
		{
			"Get all users successfully",
			h.mock,
			"SELECT uuid,name,email,phone,verified FROM users;",
			sqlmock.NewRows([]string{"some_a", "some_b", "some_c", "some_d", "some_e"}).AddRow("some-aa", "some-bb", "some-cc", "some-dd", "some-ee"),
			nil,
		},
		{
			"Something went wrong querying all users",
			h.mock,
			"SELECT uuid,name,email,phone,verified FROM users;",
			sqlmock.NewRows([]string{"some_a", "some_b", "some_c", "some_d", "some_e"}).AddRow("some-aa", "some-bb", "some-cc", "some-dd", "some-ee"),
			errors.New("some error"),
		},
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/getAllUsers", nil)

			if tt.queryErr == nil {
				tt.mock.ExpectQuery(tt.queryStr).WillReturnRows(tt.rows)
			} else {
				tt.mock.ExpectQuery(tt.queryStr).WillReturnError(tt.queryErr)
			}

			h.router.ServeHTTP(w, req)

			var actualResponse map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			if tt.queryErr == nil {
				assert.Equal(t, 200, w.Code)
				assert.NotEqual(t, "", actualResponse["resp"])

			} else {
				assert.Equal(t, 500, w.Code)
				assert.NotEqual(t, "", actualResponse["error"])
			}
		})
	}
}

func (h *HandlerTestSuite) TestLogin() {
	tests := []struct {
		name     string
		mock     sqlmock.Sqlmock
		queryStr string
		rows     *sqlmock.Rows
		queryErr error
	}{
		// {
		// 	"Sucessfull login",
		// 	h.mock,
		// 	"SELECT uuid,name,verified,psword FROM users WHERE name= ?;",
		// 	sqlmock.NewRows([]string{"some_a", "some_b", "some_c", "some_d"}).AddRow("some-aa", "some-bb", false, "some-dd"),
		// 	nil,
		// },
		{
			"Something went wrong trying to login",
			h.mock,
			"SELECT uuid,name,verified,psword FROM users WHERE name= ?;",
			sqlmock.NewRows([]string{"some_a", "some_b", "some_c", "some_d"}).AddRow("some-aa", "some-bb", false, "some-dd"),
			errors.New("some error"),
		},
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			reqParams := []byte(`{
				"name": "morpheus",
				"psword":"leader"
			}`)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqParams))

			if tt.queryErr == nil {
				tt.mock.ExpectQuery(tt.queryStr).WillReturnRows(tt.rows)
			} else {
				tt.mock.ExpectQuery(tt.queryStr).WillReturnError(tt.queryErr)
			}

			h.router.ServeHTTP(w, req)

			var actualResponse map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			if tt.queryErr == nil {
				assert.Equal(t, 200, w.Code)
				assert.NotEqual(t, "", actualResponse["msg"])
				assert.NotEqual(t, "", actualResponse["user"])
				assert.NotEqual(t, "", actualResponse["verified"])

			} else {
				assert.Equal(t, 500, w.Code)
				assert.NotEqual(t, "", actualResponse["error"])
			}
		})
	}
}

func (h *HandlerTestSuite) TestHandlerService_GetProfile() {
	tests := []struct {
		name         string
		dbMock       sqlmock.Sqlmock
		expectedCode int
		queryStr     string
		rows         *sqlmock.Rows
	}{
		{
			"Sucessfully get user profile",
			h.mock,
			200,
			"select * from links where username = ?;",
			sqlmock.NewRows([]string{"some_a", "some_b", "some_c", "some_d"}).AddRow("some-aa", "some-bb", "some-cc", "some-dd"),
		},
		{
			"something went wrong selecting user links",
			h.mock,
			500,
			"",
			nil,
		},
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/morpheus", nil)

			if tt.expectedCode == 200 {
				tt.dbMock.ExpectQuery(regexp.QuoteMeta(tt.queryStr)).WillReturnRows(tt.rows)
			}

			h.router.ServeHTTP(w, req)

			var actualResponse map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode >= 201 {
				assert.NotEqual(t, "", actualResponse["error"])
			} else {
				assert.NotEqual(t, "", actualResponse["links"])
			}
		})
	}
}
