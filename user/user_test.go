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

func (ur *fakeUserRepository) GetByUsername(username string) *User {
	if len(ur.Users) == 0 {
		return nil
	} else {
		return ur.Users[0]
	}
}

func TestNewUserService_ShouldHaveRepository(t *testing.T) {
	fakeRepo := new(fakeUserRepository)
	service := NewUserService(fakeRepo)

	assert.Equal(t, fakeRepo, service.Repo)
}

func TestCreateNewUser(t *testing.T) {
	fakeRepo := newFakeUserRepository([]*User{})
	service := NewUserService(fakeRepo)

	user := &User{Username: "Eryk"}

	err := service.Create(user)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(fakeRepo.Users))
	assert.Equal(t, user, fakeRepo.Users[0])
}

func TestCreateNewUser_When_ThereIsAlreadyOneWithSameUserName_Should_ReturnError(t *testing.T) {
	fakeRepo := newFakeUserRepository([]*User{{Username: "Eryk"}})
	service := NewUserService(fakeRepo)

	user := User{Username: "Eryk"}
	err := service.Create(&user)

	assert.NotNil(t, err)
}

func TestCreateNewUser_When_RepositoryReturnsError_Should_ReturnError(t *testing.T) {
	fakeRepo := newFakeUserRepository([]*User{})
	errorMessage := "error of fake user repository"
	fakeRepo.ErrorOnCreation = errors.New(errorMessage)

	service := NewUserService(fakeRepo)

	err := service.Create(&User{Username: "Eryk"})

	assert.NotNil(t, err)
	assert.Regexp(t, errorMessage, err.Error())
}
