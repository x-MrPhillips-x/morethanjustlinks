package handler

import (
	"database/sql"
	"testing"

	"example.com/morethanjustlinks/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type HandlerTestSuite struct {
	suite.Suite
	HandlerService *HandlerService
	db             *sql.DB
	mock           sqlmock.Sqlmock
	router         *gin.Engine
}

func (h *HandlerTestSuite) SetupTest() {
	h.db, h.mock, _ = sqlmock.New()
	h.mock.ExpectPing()

	h.HandlerService, _ = NewHandlerService(h.db, zap.NewNop().Sugar(), 3)
	h.router = h.HandlerService.SetupHandlerServiceRoutes()
}

// func TestHandlerTestSuite(t *testing.T) {
// 	suite.Run(t, new(HandlerTestSuite))
// }

func TestNewHandlerService(t *testing.T) {
	db := &mocks.DbInterface{}
	db.On("Ping").Return(nil)

	h, err := NewHandlerService(db, zap.NewNop().Sugar(), 3)

	assert.NotNil(t, h)
	assert.Nil(t, err)

	db.AssertExpectations(t)
}

// func TestNewHandlerServiceError(t *testing.T) {
// 	db := &mocks.DbInterface{}
// 	db.On("Ping").Return(errors.New("ping db timed out"))

// 	h, err := NewHandlerService(db, zap.NewNop().Sugar(), 3)

// 	assert.NotNil(t, h)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, errors.New("ping db timed out"), err)

// 	db.AssertExpectations(t)
// }

// func TestSetupHandlerServiceRoutes(t *testing.T) {
// 	db := &mocks.DbInterface{}
// 	db.On("Ping").Return(nil)

// 	h, err := NewHandlerService(db, zap.NewNop().Sugar(), 3)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, h)

// 	h.SetupHandlerServiceRoutes()

// 	db.AssertExpectations(t)
// }

// func TestGetHomeRoute(t *testing.T) {

// 	router := SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)

// 	expected, _ := json.Marshal(map[string]string{
// 		"message": "pong",
// 	})

// 	assert.Equal(t, string(expected), w.Body.String())
// }

// func TestLoginRouteInValidRequest(t *testing.T) {

// 	router := SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/login", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 400, w.Code)

// 	expected, _ := json.Marshal(map[string]string{
// 		"error": "Something went wrong with your login request...",
// 	})

// 	assert.Equal(t, string(expected), w.Body.String())
// }

// func TestLoginRouteSuccess(t *testing.T) {

// 	router := SetupRouter()
// 	w := httptest.NewRecorder()

// 	var jsonData = []byte(`{
// 		"username": "morpheus",
// 		"psword": "leader"
// 	}`)

// 	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)

// 	expected, _ := json.Marshal(map[string]string{
// 		"msg": "given username morpheus and pass leader",
// 	})

// 	assert.Equal(t, string(expected), w.Body.String())
// }

// func TestNewAccountRoute(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		reqParams    []byte
// 		expectMsgKey string
// 		expectMsg    string
// 		expectCode   int
// 	}{
// 		{
// 			"Error with input data",
// 			[]byte(`{}`),
// 			"error",
// 			"Error with input data",
// 			400,
// 		},
// 		{
// 			"Error with input data - username is not valid",
// 			[]byte(`{
// 				"name": "m",
// 				"email": "morpheus@mail.com",
// 				"phone":"7777777777",
// 				"psword":"leader"
// 			}`),
// 			"error",
// 			"Error with input data",
// 			400,
// 		},
// 		{
// 			"Error with input data - email format",
// 			[]byte(`{
// 				"name": "morpheus",
// 				"email": "morpheus",
// 				"phone":"7777777777",
// 				"psword":"leader"
// 			}`),
// 			"error",
// 			"Error with input data",
// 			400,
// 		},
// 		{
// 			"Error with input data - phone # min digits",
// 			[]byte(`{
// 				"name": "morpheus",
// 				"email": "morpheus@mail.com",
// 				"phone":"77777",
// 				"psword":"leader"
// 			}`),
// 			"error",
// 			"Error with input data",
// 			400,
// 		},
// 		{
// 			"Error with input data - phone # max digits",
// 			[]byte(`{
// 				"name": "morpheus",
// 				"email": "morpheus@mail.com",
// 				"phone":"777777777777777",
// 				"psword":"leader"
// 			}`),
// 			"error",
// 			"Error with input data",
// 			400,
// 		},
// 		{
// 			"Success creating a new account",
// 			[]byte(`{
// 				"name": "morpheus",
// 				"email": "morpheus@mail.com",
// 				"phone":"7777777777",
// 				"psword":"leader"
// 			}`),
// 			"msg",
// 			"successfully created new user",
// 			200,
// 		},
// 		{
// 			"Success creating a new account - verified true",
// 			[]byte(`{
// 				"name": "morpheus",
// 				"email": "morpheus@mail.com",
// 				"phone":"7777777777",
// 				"psword":"leader",
// 				"verified":true
// 			}`),
// 			"msg",
// 			"successfully created new user",
// 			200,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()
// 			w := httptest.NewRecorder()

// 			req, _ := http.NewRequest("POST", "/newAccount", bytes.NewBuffer(tt.reqParams))

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tt.expectCode, w.Code)

// 			var actualResponse map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &actualResponse)

// 			if tt.expectMsgKey == "error" {
// 				assert.Equal(t, tt.expectMsg, actualResponse["error"])
// 			} else {
// 				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
// 			}
// 		})
// 	}
// }

// func TestGetAllUsersRoute(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		reqParams    []byte
// 		expectMsgKey string
// 		expectMsg    string
// 		expectCode   int
// 	}{
// 		{
// 			"Success fetching all users",
// 			nil,
// 			"msg",
// 			"Successfully fetched all users",
// 			200,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()
// 			w := httptest.NewRecorder()

// 			req, _ := http.NewRequest("GET", "/getAllUsers", bytes.NewBuffer(tt.reqParams))

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tt.expectCode, w.Code)

// 			var actualResponse map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &actualResponse)

// 			if tt.expectMsgKey == "error" {
// 				assert.Equal(t, tt.expectMsg, actualResponse["error"])
// 			} else {
// 				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
// 			}
// 		})
// 	}
// }

// func TestUpdateUserRoute(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		reqParams    []byte
// 		expectMsgKey string
// 		expectMsg    string
// 		expectCode   int
// 	}{
// 		{
// 			"Success update user",
// 			[]byte(`{
// 				"uuid":"ba198f66-7453-4ee5-8a0c-8e693d408658",
// 				"name": "nameBeenUpdatedBuddy",
// 				"email": "nameBeenUpdatedBuddy@mail.com",
// 				"phone":"8888888888",
// 				"psword":"nobody",
// 				"verified":true
// 			}`),
// 			"msg",
// 			"Successfully updated user",
// 			200,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()
// 			w := httptest.NewRecorder()

// 			req, _ := http.NewRequest("POST", "/update", bytes.NewBuffer(tt.reqParams))

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tt.expectCode, w.Code)

// 			var actualResponse map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &actualResponse)

// 			if tt.expectMsgKey == "error" {
// 				assert.Equal(t, tt.expectMsg, actualResponse["error"])
// 			} else {
// 				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
// 			}
// 		})
// 	}
// }

// func TestDeleteUserRoute(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		reqParams    []byte
// 		expectMsgKey string
// 		expectMsg    string
// 		expectCode   int
// 	}{
// 		{
// 			"Success delete user",
// 			[]byte(`{
// 				"name": "nameBeenUpdatedBuddy"
// 			}`),
// 			"msg",
// 			"Successfully deleted user",
// 			200,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()
// 			w := httptest.NewRecorder()

// 			req, _ := http.NewRequest("POST", "/delete", bytes.NewBuffer(tt.reqParams))

// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tt.expectCode, w.Code)

// 			var actualResponse map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &actualResponse)

// 			if tt.expectMsgKey == "error" {
// 				assert.Equal(t, tt.expectMsg, actualResponse["error"])
// 			} else {
// 				assert.Equal(t, tt.expectMsg, actualResponse["msg"])
// 			}
// 		})
// 	}
// }
