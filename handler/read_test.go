package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (h *HandlerTestSuite) TestGetAllUsers() {
	tests := []struct {
		name         string
		reqParams    []byte
		expectMsgKey string
		expectMsg    string
		expectCode   int
		queryErr     error
		adaptErr     error
	}{
		// {
		// 	"Error querying db for all users",
		// 	[]byte(`{}`),
		// 	"error",
		// 	"error fetching all users",
		// 	500,
		// 	errors.New("some error"),
		// 	nil,
		// },
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			var sqlRows *sql.Rows
			h.db_mock.On(
				"Query",
				SELECT_ALL_USERS).Return(sqlRows, tt.queryErr).Once()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", "/getAllUsers", bytes.NewBuffer(tt.reqParams))

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
