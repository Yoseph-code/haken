package db

import (
	"os"
	"path"

	safeMap "github.com/Yoseph-code/haken/internal/safe_map"
)

const (
	defaultSourceName string = "default"
	extFile           string = ".db"
	defaultDir        string = "haken"
)

type DB struct {
	SourceName string

	sm *safeMap.SafeMap[string]
}

func New(sourceName ...string) *DB {
	s := defaultSourceName

	if len(sourceName) > 0 {
		s = sourceName[0]
	}

	return &DB{
		SourceName: s,
		sm:         safeMap.NewSafeMap[string](),
	}
}

func (db *DB) GetSourceDB() (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(pwd, defaultDir, db.SourceName+extFile), nil
}
