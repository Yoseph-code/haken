package db

import (
	"github.com/Yoseph-code/haken/internal/fs"
)

func (db *DB) Set(key, value string) error {
	filename, err := db.GetSourceDB()

	if err != nil {
		return err
	}

	return fs.Append(filename, map[string]string{
		key: value,
	})
}

func (db *DB) Get(key string) (string, bool) {
	filename, err := db.GetSourceDB()

	if err != nil {
		return err.Error(), false
	}

	data, err := fs.Load(filename)

	if err != nil {
		return err.Error(), false
	}

	val, ok := data[key]

	if !ok {
		return "no value finded for this key", false
	}

	return val, true
}
