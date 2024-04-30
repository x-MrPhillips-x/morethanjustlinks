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

func (h *HandlerTestSuite) TestNewAccount() {
	tests := []struct {
		name       string
		params     []byte
		prepare    func()
		statusCode int
		msg        string
	}{
		{
			name:       "Request with invalid params",
			params:     []byte(`{"foo":"bar}`),
			statusCode: http.StatusBadRequest,
		},
		{
			name:   "failed to create account because user already exists",
			params: []byte(`{"name":"morpheusss","email":"morpheusss@mail.com","phone":"6155577777","psword":"secret","role":"admin"}`),
			prepare: func() {
				expectedSQL := ".+"
				addRow := sqlmock.NewRows([]string{
					"uuid", "name", "email", "phone", "psword", "verified", "role",
				}).AddRow(
					"ff5c13a2-d04b-44ef-a337-6b45b1a8dd4c",
					"morpheusss",
					"morpheusss@mail.com",
					"6155577777",
					"$2a$14$0u.ZyqgZTyZyPN63JlVdz.SsdxBgfpLhDAE14oL4iTxhp.a..f/dC",
					"false",
					"admin",
				)

				h.mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
			},
			statusCode: http.StatusOK,
			msg:        "email or phone number entered is already in use",
		},
		{
			name:   "successfully create a new admin account",
			params: []byte(`{"name":"morpheusss","email":"morpheusss@mail.com","phone":"6155577777","psword":"secret","role":"admin"}`),
			prepare: func() {
				expectedSQL := ".+"
				h.mock.ExpectQuery(expectedSQL).WillReturnRows(&sqlmock.Rows{})

				h.mock.ExpectBegin()
				h.mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(0, 1))
				h.mock.ExpectCommit()
			},
			statusCode: http.StatusOK,
			msg:        "successfully created morpheusss",
		},
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			if tt.prepare != nil {
				tt.prepare()
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/newAccount", bytes.NewBuffer(tt.params))
			h.router.ServeHTTP(w, req)

			var resp map[string]string
			json.Unmarshal(w.Body.Bytes(), &resp)

			assert.Equal(h.T(), tt.statusCode, w.Code)
			assert.Equal(h.T(), tt.msg, resp["msg"])
			assert.Nil(h.T(), h.mock.ExpectationsWereMet())
		})
	}
}
