package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (h *HandlerTestSuite) TestHandlerService_DeleteLink() {

	tests := []struct {
		name         string
		reqParams    DeleteLinkRequest
		expectCode   int
		expectMsgKey string
		expectMsg    string
		dbMocks      sqlmock.Sqlmock
	}{
		{
			"missing required uuid",
			DeleteLinkRequest{
				UUID: "",
			},
			400,
			"error",
			"error missing required uuid",
			h.mock,
		},
		{
			"this is the first test",
			DeleteLinkRequest{
				UUID: "30a1ce10-e885-4652-a9cc-8c2bff55f8f2",
			},
			200,
			"msg",
			"Successfully deleted link",
			h.mock,
		},
	}
	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()

			body, _ := json.Marshal(tt.reqParams)
			req, _ := http.NewRequest("POST", "/deleteLink", bytes.NewBuffer(body))

			if tt.expectCode == 200 {
				tt.dbMocks.ExpectExec("DELETE FROM links WHERE uuid = ?").WillReturnResult(sqlmock.NewResult(1, 1))
			}

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

func (h *HandlerTestSuite) TestHandlerService_DeleteUser() {
	tests := []struct {
		name         string
		reqParams    DeleteUserRequest
		expectCode   int
		expectMsgKey string
		expectMsg    string
		dbMocks      sqlmock.Sqlmock
	}{
		{
			"missing required uuid",
			DeleteUserRequest{
				Name: "",
			},
			400,
			"error",
			"error missing required name",
			h.mock,
		},
		{
			"Success user name is deleted",
			DeleteUserRequest{
				Name: "superman",
			},
			200,
			"msg",
			"Successfully deleted user",
			h.mock,
		},
	}
	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()

			body, _ := json.Marshal(tt.reqParams)
			req, _ := http.NewRequest("POST", "/deleteUser", bytes.NewBuffer(body))

			if tt.expectCode == 200 {
				tt.dbMocks.ExpectExec("DELETE FROM users WHERE name = ?").WillReturnResult(sqlmock.NewResult(1, 1))
			}

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
