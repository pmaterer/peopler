package controller

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pmaterer/peopler/user"
	"github.com/stretchr/testify/assert"
)

var (
	userPayload          = `{"firstName":"Shane","lastName":"Glass"}`
	userPayloadMalformed = `{"firstName":"Shane","lastName":"Glass"`

	updateUserPayload          = `{"firstName":"Shane","lastName":"Glas"}`
	updateUserPayloadMalformed = `{"firstName":"Shane","lastName":"Glas"`

	testUser = user.User{
		ID:        1,
		FirstName: "Shane",
		LastName:  "Glass",
	}
	testUserPayload = `{"id":1,"firstName":"Shane","lastName":"Glass"}`

	testUsers = []user.User{
		{
			ID:        2,
			FirstName: "Stephen",
			LastName:  "King",
		},
		{
			ID:        3,
			FirstName: "Herman",
			LastName:  "Melville",
		},
		{
			ID:        4,
			FirstName: "Stanley",
			LastName:  "Kubrick",
		},
	}
	testUsersPayload = `[{"id":2,"firstName":"Stephen","lastName":"King"},{"id":3,"firstName":"Herman","lastName":"Melville"},{"id":4,"firstName":"Stanley","lastName":"Kubrick"}]`
)

type mockService struct {
	CreateUserFunc  func(u user.User) error
	GetAllUsersFunc func() ([]user.User, error)
	GetUserFunc     func(id int64) (user.User, error)
	UpdateUserFunc  func(u user.User) error
	DeleteUserFunc  func(id int64) error
}

func (s *mockService) CreateUser(u user.User) error        { return s.CreateUserFunc(u) }
func (s *mockService) GetAllUsers() ([]user.User, error)   { return s.GetAllUsersFunc() }
func (s *mockService) GetUser(id int64) (user.User, error) { return s.GetUserFunc(id) }
func (s *mockService) UpdateUser(u user.User) error        { return s.UpdateUserFunc(u) }
func (s *mockService) DeleteUser(id int64) error           { return s.DeleteUserFunc(id) }

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(u user.User) error
		payload     string
	}{
		{
			name:        "Create user OK",
			errExpected: false,
			method: func(u user.User) error {
				return nil
			},
			payload: userPayload,
		},
		{
			name:        "Create user error",
			errExpected: true,
			method: func(u user.User) error {
				return errors.New("bad stuff")
			},
			payload: userPayload,
		},
		{
			name:        "Create user error",
			errExpected: true,
			method: func(u user.User) error {
				return nil
			},
			payload: userPayloadMalformed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				CreateUserFunc: tt.method,
			}
			c := NewController(s)

			req, err := http.NewRequest("POST", "/user", strings.NewReader(tt.payload))
			assert.Nil(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(c.CreateUser())

			handler.ServeHTTP(rr, req)

			if tt.errExpected {
				if tt.payload == userPayloadMalformed {
					assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
				} else {
					assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
				}
			} else {
				assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func() ([]user.User, error)
	}{
		{
			name:        "Get all users OK",
			errExpected: false,
			method: func() ([]user.User, error) {
				return testUsers, nil
			},
		},
		{
			name:        "Get all users error",
			errExpected: true,
			method: func() ([]user.User, error) {
				return testUsers, errors.New("bad stuff")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				GetAllUsersFunc: tt.method,
			}
			c := NewController(s)

			req, err := http.NewRequest("GET", "/users", nil)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(c.GetAllUsers())

			handler.ServeHTTP(rr, req)

			if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
			} else {
				respBody := rr.Result().Body
				defer respBody.Close()
				payload, _ := ioutil.ReadAll(respBody)

				assert.Equal(t, testUsersPayload, string(payload))
				assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(id int64) (user.User, error)
	}{
		{
			name:        "Get user OK",
			errExpected: false,
			method: func(id int64) (user.User, error) {
				return testUser, nil
			},
		},
		{
			name:        "Get user error",
			errExpected: true,
			method: func(id int64) (user.User, error) {
				return testUser, errors.New("bad stuff")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				GetUserFunc: tt.method,
			}
			c := NewController(s)

			req, err := http.NewRequest("GET", "/user/1", nil)
			assert.Nil(t, err)
			rr := httptest.NewRecorder()
			vars := map[string]string{
				"id": "1",
			}
			req = mux.SetURLVars(req, vars)

			handler := http.HandlerFunc(c.GetUser())

			handler.ServeHTTP(rr, req)

			if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
			} else {
				respBody := rr.Result().Body
				defer respBody.Close()
				payload, _ := ioutil.ReadAll(respBody)

				assert.Equal(t, testUserPayload, string(payload))
				assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(u user.User) error
		payload     string
	}{
		{
			name:        "Update user OK",
			errExpected: false,
			method: func(u user.User) error {
				return nil
			},
			payload: updateUserPayload,
		},
		{
			name:        "Update user error",
			errExpected: true,
			method: func(u user.User) error {
				return errors.New("bad stuff")
			},
			payload: updateUserPayload,
		},
		{
			name:        "Update user malformed payload error",
			errExpected: true,
			method: func(u user.User) error {
				return errors.New("bad stuff")
			},
			payload: updateUserPayloadMalformed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				UpdateUserFunc: tt.method,
			}
			c := NewController(s)

			req, err := http.NewRequest("PUT", "/user/1", strings.NewReader(tt.payload))
			assert.Nil(t, err)

			vars := map[string]string{
				"id": "1",
			}
			req = mux.SetURLVars(req, vars)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(c.UpdateUser())

			handler.ServeHTTP(rr, req)

			if tt.errExpected {
				if tt.payload == updateUserPayloadMalformed {
					assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
				} else {
					assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
				}
			} else {
				assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(id int64) error
	}{
		{
			name:        "Delete user OK",
			errExpected: false,
			method: func(id int64) error {
				return nil
			},
		},
		{
			name:        "Delete user error",
			errExpected: true,
			method: func(id int64) error {
				return errors.New("bad stuff")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				DeleteUserFunc: tt.method,
			}
			c := NewController(s)

			req, err := http.NewRequest("DELETE", "/user/1", nil)
			assert.Nil(t, err)

			vars := map[string]string{
				"id": "1",
			}
			req = mux.SetURLVars(req, vars)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(c.DeleteUser())

			handler.ServeHTTP(rr, req)

			if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
			}
		})
	}
}
