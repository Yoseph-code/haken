package db

import (
	"os"
	"path"
)

const (
	defaultSourceName string = "default"
	extFile           string = ".db"
	defaultDir        string = "haken"
)

type DB struct {
	SourceName string
}

func New(sourceName ...string) *DB {
	s := defaultSourceName

	if len(sourceName) > 0 {
		s = sourceName[0]
	}

	return &DB{
		SourceName: s,
	}
}

func (db *DB) GetSourceDB() (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(pwd, defaultDir, db.SourceName+extFile), nil
}
