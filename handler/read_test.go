package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (h *HandlerTestSuite) TestGetAllUsers() {
	sqlRows := &sql.Rows{}
	tests := []struct {
		name         string
		reqParams    []byte
		expectMsgKey string
		expectMsg    string
		expectCode   int
		adaptErr     error
		mocks        func()
	}{
		{
			"Error querying db for all users",
			[]byte(`{}`),
			"error",
			"something went wrong...",
			500,
			nil,
			func() {
				h.db_mock.On(
					"Query",
					SELECT_ALL_USERS).Return(sqlRows, errors.New("some error"))
			},
		},
		// {
		// 	"Error rows next returns false",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"something went wrong...",
		// 	500,
		// 	nil,
		// 	func() {
		// 		h.db_mock.On(
		// 			"Query",
		// 			SELECT_ALL_USERS).Return(sqlRows, nil)
		// 		h.rows_mock.On("Next").Return(false)
		// 	},
		// },
		// {
		// 	"Error rows scan",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"something went wrong...",
		// 	500,
		// 	nil,
		// 	func() {
		// 		h.db_mock.On(
		// 			"Query",
		// 			SELECT_ALL_USERS).Return(sqlRows, nil)
		// 		h.rows_mock.On("Next").Return(true)
		// 		h.rows_mock.On("Scan").Return(errors.New("some error"))
		// 	},
		// },
		// {
		// 	"Success get all users",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"something went wrong...",
		// 	200,
		// 	nil,
		// 	func() {
		// 		h.db_mock.On(
		// 			"Query",
		// 			SELECT_ALL_USERS).Return(
		// 			sqlmock.NewRows([]string{"uuid", "name", "email", "phone", "verified"}).
		// 				AddRow(
		// 					"08eb0d36-5669-4e6c-b4a3-6097bddd75bc",
		// 					"ham",
		// 					"burger@mail.com",
		// 					"8888888888",
		// 					true,
		// 				), nil,
		// 		)
		// 		h.rows_mock.On("Next").Return(true)
		// 		h.rows_mock.On(
		// 			"Scan",
		// 			mock.Anything,
		// 			mock.Anything,
		// 			mock.Anything,
		// 			mock.Anything,
		// 			mock.Anything,
		// 		).Return(nil)
		// 	},
		// },
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()

			tt.mocks()

			req, _ := http.NewRequest("GET", "/getAllUsers", bytes.NewBuffer(tt.reqParams))

			h.router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			var actualResponse map[string]interface{}

			fmt.Println("This is the response:", w.Body.String())
			json.Unmarshal(w.Body.Bytes(), &actualResponse)

			if tt.expectMsgKey == "error" {
				assert.Equal(t, tt.expectMsg, actualResponse["error"])
			} else {
				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
			}
		})
	}
}
