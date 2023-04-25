package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakeUserRepository struct {
	Users           []*User
	ErrorOnCreation error
}

func newFakeUserRepository(users []*User) *fakeUserRepository {
	return &fakeUserRepository{Users: users}
}

func (ur *fakeUserRepository) Create(user *User) error {
	if ur.ErrorOnCreation != nil {
		return ur.ErrorOnCreation
	}

	ur.Users = append(ur.Users, user)
	return nil
}

func (ur *fakeUserRepository) GetByUsername(username string) (*User, error) {
	if len(ur.Users) == 0 {
		return nil, nil
	} else {
		return ur.Users[0], nil
	}
}

func (ur *fakeUserRepository) List() ([]*User, error) {
	panic("not implemented")
}

func TestNewUserService_ShouldHaveRepository(t *testing.T) {
	fakeRepo := new(fakeUserRepository)
	service := NewService(fakeRepo)

	assert.Equal(t, fakeRepo, service.repo)
}

func TestCreateNewUser(t *testing.T) {
	fakeRepo := newFakeUserRepository([]*User{})
	service := NewService(fakeRepo)

	user := &User{Username: "Eryk"}

	err := service.Create(user)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(fakeRepo.Users))
	assert.Equal(t, user, fakeRepo.Users[0])
}

func TestCreateNewUser_When_ThereIsAlreadyOneWithSameUserName_Should_ReturnError(t *testing.T) {
	fakeRepo := newFakeUserRepository([]*User{{Username: "Eryk"}})
	service := NewService(fakeRepo)

	user := User{Username: "Eryk"}
	err := service.Create(&user)

	assert.NotNil(t, err)
}

func TestCreateNewUser_When_RepositoryReturnsError_Should_ReturnError(t *testing.T) {
	fakeRepo := newFakeUserRepository([]*User{})
	errorMessage := "error of fake user repository"
	fakeRepo.ErrorOnCreation = errors.New(errorMessage)

	service := NewService(fakeRepo)

	err := service.Create(&User{Username: "Eryk"})

	assert.NotNil(t, err)
	assert.Regexp(t, errorMessage, err.Error())
}
