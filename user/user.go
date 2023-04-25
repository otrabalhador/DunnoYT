package user

import (
	"fmt"
)

type User struct {
	Username string
}

type Repository interface {
	Create(*User) error
	List() ([]*User, error)
	GetByUsername(username string) (*User, error)
}

type Service struct {
	repo Repository
}

func NewService(userRepo Repository) *Service {
	return &Service{
		repo: userRepo,
	}
}

func (s *Service) Create(user *User) error {
	alreadyExistent, err := s.repo.GetByUsername(user.Username)
	if err != nil {
		return fmt.Errorf("Error getting user by username: %w", err)
	}

	if alreadyExistent != nil {
		return fmt.Errorf("username %v already exists", user.Username)
	}

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (s *Service) List() ([]*User, error) {
	return s.repo.List()
}

func (s *Service) Get(username string) (*User, error) {
	return s.repo.GetByUsername(username)
}
