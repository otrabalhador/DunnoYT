package user

import (
	"fmt"
)

type User struct {
	Username string
}

type Repository interface {
	Create(*User) error
	GetByUsername(username string) *User
}

type Service struct {
	Repo Repository
}

func NewUserService(userRepo Repository) *Service {
	return &Service{
		Repo: userRepo,
	}
}

func (s *Service) Create(user *User) error {
	alreadyExistent := s.Repo.GetByUsername(user.Username)
	if alreadyExistent != nil {
		return fmt.Errorf("username %v already exists", user.Username)
	}

	if err := s.Repo.Create(user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
