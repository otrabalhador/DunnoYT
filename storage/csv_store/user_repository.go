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
	defer file.Close()

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

func (c *CsvUserRepository) GetByUsername(username string) *domainUser.User {
	//TODO implement me
	panic("implement me")
}

func createFileIfNotExists(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		return nil // File already exists
	}

	if _, err = os.Create(filePath); err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	return nil
}
