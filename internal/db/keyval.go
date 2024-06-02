package db

import (
	"github.com/Yoseph-code/haken/internal/fs"
)

func (db *DB) Set(key, value string) error {
	return fs.Append(db.sourceName, map[string]string{
		key: value,
	})
}

func (db *DB) Get(key string) (string, bool) {
	data, err := fs.Load(db.sourceName)

	if err != nil {
		return err.Error(), false
	}

	val, ok := data[key]

	if !ok {
		return "no value finded for this key", false
	}

	return val, true
}
