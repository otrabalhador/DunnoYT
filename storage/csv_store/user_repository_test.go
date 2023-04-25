package csv_store

import (
	domainUser "DunnoYT/user"
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TODO: Clean up csv files on test teardown

func TestNewCsvUserRepository_Should_CreateFileIfNotExists(t *testing.T) {
	fileName := "test_constructor.csv"
	os.Remove(fileName)

	_, err := NewCsvUserRepository(fileName)
	assert.Nil(t, err)

	f, err := os.Stat(fileName)

	assert.Nil(t, err)
	assert.NotNil(t, f)
}

func TestNewCsvUserRepository_Should_NotOverwriteFileIfAlreadyExists(t *testing.T) {
	// TODO
}

func TestCreate_Should_WriteAsCsv(t *testing.T) {
	fileName := "test_create_once.csv"
	os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName)

	user := &domainUser.User{Username: "John"}
	err := repo.Create(user)

	assert.Nil(t, err)

	file, _ := os.Open(fileName)
	defer file.Close()

	lines, _ := csv.NewReader(file).ReadAll()

	assert.Equal(t, 1, len(lines))
	assert.Equal(t, user.Username, lines[0][0])
}

func TestCreateTwice_Should_BreakLineAndAddSecondUser(t *testing.T) {
	fileName := "test_user_create_twice.csv"
	os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName)

	user1 := &domainUser.User{Username: "John"}
	user2 := &domainUser.User{Username: "Paul"}
	_ = repo.Create(user1)
	_ = repo.Create(user2)

	file, _ := os.Open(fileName)
	defer file.Close()

	lines, _ := csv.NewReader(file).ReadAll()

	assert.Equal(t, 2, len(lines))
	assert.Equal(t, user1.Username, lines[0][0])
	assert.Equal(t, user2.Username, lines[1][0])
}
