package service

import (
	"log"

	"github.com/pmaterer/peopler/user"
)

type repository interface {
	CreateUser(u user.User) (int64, error)
	GetAllUsers() ([]user.User, error)
	GetUser(id int64) (user.User, error)
	UpdateUser(u user.User) (int64, error)
	DeleteUser(id int64) (int64, error)
}

type Service struct {
	repository repository
}

func NewService(r repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) CreateUser(u user.User) error {
	id, err := s.repository.CreateUser(u)
	if err != nil {
		return err
	}
	log.Printf("Created new user #%d\n", id)
	return nil
}

func (s *Service) GetUser(id int64) (user.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *Service) GetAllUsers() ([]user.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *Service) UpdateUser(u user.User) error {
	id, err := s.repository.UpdateUser(u)
	if err != nil {
		return err
	}
	log.Printf("Updated user #%d\n", id)
	return nil
}

func (s *Service) DeleteUser(id int64) error {
	id, err := s.repository.DeleteUser(id)
	if err != nil {
		return err
	}
	log.Printf("Deleted user #%d\n", id)
	return nil
}
