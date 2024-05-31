package db

import (
	"fmt"
	"io/fs"
	"os"
)

const (
	mainPath = "haken"
	ext      = "db"
)

func (db *DB) CreateOrReadFile(pwd string) (string, error) {
	p := fmt.Sprintf("%s/%s", pwd, mainPath)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.Mkdir(mainPath, fs.FileMode(os.O_CREATE|os.O_RDWR|os.O_APPEND))

		if err != nil {
			return "", err
		}
	}

	return p, nil
}

func (db *DB) Init() error {
	pwd, err := os.Getwd()

	if err != nil {
		return err
	}

	dir, err := db.CreateOrReadFile(pwd)

	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s.%s", dir, db.Config.FileName, ext)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)

		if err != nil {
			return err
		}

		defer file.Close()

		err = file.Chmod(os.FileMode(os.O_APPEND | os.O_CREATE | os.O_RDWR))

		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	db.db = &fileName

	db.Read()

	return nil
}
