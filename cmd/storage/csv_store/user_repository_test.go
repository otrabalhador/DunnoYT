package csv_store

import (
	domainUser "DunnoYT/cmd/user"
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestNewCsvUserRepository_Should_CreateFileIfNotExists(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	_, err := NewCsvUserRepository(fileName, false)
	assert.Nil(t, err)

	f, err := os.Stat(fileName)

	assert.Nil(t, err)
	assert.NotNil(t, f)
}

func TestNewCsvUserRepositoryWithClearPreviousContentTrue_Should_RemoveFileAndCreateAnother(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()
	_, _ = f.WriteString("John\nEryk\nJorge")

	_, err := NewCsvUserRepository(fileName, true)
	assert.Nil(t, err)

	file, _ := os.Open(fileName)
	defer file.Close()

	bytes := make([]byte, 10)
	readBytes, _ := file.Read(bytes)
	assert.Equal(t, 0, readBytes)
}

func TestCreate_Should_WriteAsCsv(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName, false)

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
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName, false)

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

func TestList_When_ThereIsNone_Should_ReturnEmptyArray(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName, false)

	users, err := repo.List()

	assert.Nil(t, err)
	assert.Empty(t, users)
}

func TestList_When_ThereAreUsers_Should_ReturnAllUsers(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName, false)

	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()
	_, err := f.WriteString("John\nEryk\nJorge")

	expected := []domainUser.User{
		{Username: "John"},
		{Username: "Eryk"},
		{Username: "Jorge"},
	}

	users, err := repo.List()

	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(users))

	for i := 0; i < len(users); i++ {
		assert.Equal(t, expected[i].Username, users[i].Username)
	}
}

func TestGetByUsername_When_NotFound_Should_ReturnNil(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName, false)
	_ = repo.Create(&domainUser.User{Username: "John"})

	user, err := repo.GetByUsername("Eryk")

	assert.Nil(t, err)
	assert.Nil(t, user)
}

func TestGetByUsername_When_Found_Should_ReturnUser(t *testing.T) {
	fileName := GetNewFileName()
	defer os.Remove(fileName)

	repo, _ := NewCsvUserRepository(fileName, false)
	_ = repo.Create(&domainUser.User{Username: "Eryk"})

	user, err := repo.GetByUsername("Eryk")

	assert.Nil(t, err)
	assert.Equal(t, "Eryk", user.Username)
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetNewFileName() string {
	const charset = "" +
		"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b) + ".csv"
}
