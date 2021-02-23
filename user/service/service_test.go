package service

import (
	"errors"
	"testing"

	"github.com/pmaterer/peopler/user"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	CreateUserFunc  func(u user.User) (int64, error)
	GetAllUsersFunc func() ([]user.User, error)
	GetUserFunc     func(id int64) (user.User, error)
	UpdateUserFunc  func(u user.User) (int64, error)
	DeleteUserFunc  func(id int64) (int64, error)
}

func (r *mockRepository) CreateUser(u user.User) (int64, error) { return r.CreateUserFunc(u) }
func (r *mockRepository) GetAllUsers() ([]user.User, error)     { return r.GetAllUsersFunc() }
func (r *mockRepository) GetUser(id int64) (user.User, error)   { return r.GetUserFunc(id) }
func (r *mockRepository) UpdateUser(u user.User) (int64, error) { return r.UpdateUserFunc(u) }
func (r *mockRepository) DeleteUser(id int64) (int64, error)    { return r.DeleteUserFunc(id) }

var (
	testUser = user.User{
		ID:        1,
		FirstName: "Stephen",
		LastName:  "King",
	}

	testUsers = []user.User{
		{
			ID:        2,
			FirstName: "Herman",
			LastName:  "Melville",
		},
		{
			ID:        3,
			FirstName: "Haruki",
			LastName:  "Murakami",
		},
		{
			ID:        4,
			FirstName: "Stanley",
			LastName:  "Kubrick",
		},
	}
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(u user.User) (int64, error)
	}{
		{
			name:        "Create user OK",
			errExpected: false,
			method: func(u user.User) (int64, error) {
				return testUser.ID, nil
			},
		},
		{
			name:        "Create user error",
			errExpected: true,
			method: func(u user.User) (int64, error) {
				return testUser.ID, errors.New("bad things")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{CreateUserFunc: tt.method}
			s := NewService(r)
			err := s.CreateUser(testUser)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
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
				return testUsers, errors.New("bad things")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{GetAllUsersFunc: tt.method}
			s := NewService(r)
			users, err := s.GetAllUsers()
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				for i, user := range users {
					assert.Equal(t, testUsers[i].ID, user.ID)
					assert.Equal(t, testUsers[i].FirstName, user.FirstName)
					assert.Equal(t, testUsers[i].LastName, user.LastName)
				}
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
				return testUser, errors.New("bad things")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{GetUserFunc: tt.method}
			s := NewService(r)
			user, err := s.GetUser(testUser.ID)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testUser.ID, user.ID)
				assert.Equal(t, testUser.FirstName, user.FirstName)
				assert.Equal(t, testUser.LastName, user.LastName)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(u user.User) (int64, error)
	}{
		{
			name:        "Update user OK",
			errExpected: false,
			method: func(u user.User) (int64, error) {
				return testUser.ID, nil
			},
		},
		{
			name:        "Update user error",
			errExpected: true,
			method: func(u user.User) (int64, error) {
				return testUser.ID, errors.New("bad things")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{UpdateUserFunc: tt.method}
			s := NewService(r)
			err := s.UpdateUser(testUser)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(id int64) (int64, error)
	}{
		{
			name:        "Delete user OK",
			errExpected: false,
			method: func(id int64) (int64, error) {
				return testUser.ID, nil
			},
		},
		{
			name:        "Delete user error",
			errExpected: true,
			method: func(id int64) (int64, error) {
				return testUser.ID, errors.New("bad things")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{DeleteUserFunc: tt.method}
			s := NewService(r)
			err := s.DeleteUser(testUser.ID)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
