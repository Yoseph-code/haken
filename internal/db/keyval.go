package db

import (
	"strings"

	"github.com/Yoseph-code/haken/internal/fs"
)

func (db *DB) Set(key string, value []byte) error {
	filename, err := db.GetSourceDB()

	if err != nil {
		return err
	}

	return fs.Append(filename, map[string]string{
		key: strings.TrimSuffix(string(value), "\n"),
	})
}

func (db *DB) Get(key string) (string, bool) {
	filename, err := db.GetSourceDB()

	if err != nil {
		return "", false
	}

	data, err := fs.Load(filename)

	if err != nil {
		return "", false
	}

	if val, ok := data[key]; ok {
		return val, true
	}

	return "", false
}
