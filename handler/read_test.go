package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/morethanjustlinks/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (h *HandlerTestSuite) TestGetAllUsers() {
	tests := []struct {
		name    string
		params  []byte
		prepare func()
		resp    []models.User
	}{
		{
			name:   "no users found",
			params: []byte(`{}`),
			prepare: func() {
				expectedSQL := ".+"
				h.mock.ExpectQuery(expectedSQL).WillReturnRows(&sqlmock.Rows{})
			},
			resp: []models.User{},
		},
		{
			name:   "successfully get all users",
			params: []byte(`{}`),
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
			resp: []models.User{
				{
					// ID:       "ff5c13a2-d04b-44ef-a337-6b45b1a8dd4c",
					Name:     "morpheusss",
					Email:    "morpheusss@mail.com",
					Phone:    "6155577777",
					Psword:   "$2a$14$0u.ZyqgZTyZyPN63JlVdz.SsdxBgfpLhDAE14oL4iTxhp.a..f/dC",
					Verified: false,
					Role:     "admin",
				},
			},
		},
	}

	for _, tt := range tests {
		h.T().Run(tt.name, func(t *testing.T) {

			if tt.prepare != nil {
				tt.prepare()
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/getAllUsers", bytes.NewBuffer(tt.params))

			h.router.ServeHTTP(w, req)

			var resp []models.User

			json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Equal(h.T(), len(tt.resp), len(resp))
			assert.Nil(h.T(), h.mock.ExpectationsWereMet())
		})

	}

}
