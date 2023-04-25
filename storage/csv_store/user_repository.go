package csv_store

import (
	domainUser "DunnoYT/user"
	"encoding/csv"
	"fmt"
	"os"
)

type CsvUserRepository struct {
	FilePath string
}

func NewCsvUserRepository(filePath string) (*CsvUserRepository, error) {
	err := createFileIfNotExists(filePath)
	if err != nil {
		return nil, err
	}

	return &CsvUserRepository{FilePath: filePath}, nil
}

func (c *CsvUserRepository) Create(user *domainUser.User) error {
	file, err := os.OpenFile(c.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing")
		}
	}(file)

	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{user.Username})
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func (c *CsvUserRepository) List() ([]*domainUser.User, error) {
	file, err := os.Open(c.FilePath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error opening csv: %w", err)
	}

	users := []*domainUser.User{}
	for _, line := range lines {
		users = append(users, &domainUser.User{
			Username: line[0],
		})
	}

	return users, nil
}

func (c *CsvUserRepository) GetByUsername(username string) (*domainUser.User, error) {
	users, err := c.List()
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	for _, user := range users {
		if username == user.Username {
			return user, nil
		}
	}

	return nil, nil
}

func createFileIfNotExists(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		return nil // File already exists
	}

	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	return nil
}
