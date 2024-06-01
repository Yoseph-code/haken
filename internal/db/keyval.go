package db

import (
	"strings"
	"unicode"

	"github.com/Yoseph-code/haken/internal/fs"
)

func (db *DB) Set(key string, value []byte) error {
	filename, err := db.GetSourceDB()

	if err != nil {
		return err
	}

	var data map[string]interface{}

	err = fs.LoadFromFile(filename, &data)

	if err != nil {
		return err
	}

	data[key] = strings.TrimRightFunc(string(value), unicode.IsSpace)

	if err := fs.SaveToFile(
		filename,
		map[string]interface{}{
			key: strings.TrimRightFunc(string(value), unicode.IsSpace),
		},
	); err != nil {
		return err
	}

	return nil
}

func (db *DB) Get(key string) (string, bool) {
	filename, err := db.GetSourceDB()

	if err != nil {
		return "", false
	}

	var data map[string]interface{}

	err = fs.LoadFromFile(filename, &data)

	if err != nil {
		return "", false
	}

	if val, ok := data[key]; ok {
		return val.(string), true
	}

	return "", false
}
