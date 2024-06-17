package db

import (
	"errors"
	"os"
)

func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

func CreateDB(filename string) error {
	if ok := IsFileExists(filename); ok {
		return errors.New("file already exists")
	}

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	return file.Close()
}
