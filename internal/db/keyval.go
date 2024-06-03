package db

import (
	"errors"

	"github.com/Yoseph-code/haken/internal/fs"
)

func (db *DB) Set(key, value string) error {
	return fs.Append(db.sourceName, map[string]string{
		key: value,
	})
}

func (db *DB) Get(key string) (string, bool) {
	// data, err := fs.Load(db.sourceName)

	// if err != nil {
	// 	return err.Error(), false
	// }

	// // val, ok := data[key]

	// // if !ok {
	// // 	return "no value finded for this key", false
	// // }

	// return val, true

	return "", false
}

func (db *DB) Read(key string) (string, error) {
	data, err := fs.Load(db.sourceName)

	if err != nil {
		return "", err
	}

	val, ok := data.Get(key)

	if !ok {
		return "", errors.New("no value finded for this key")
	}

	return val, nil
}

// func (db *DB) Delete(key string) error {
// 	data, err := fs.Load(db.sourceName)

// 	if err != nil {
// 		return err
// 	}

// 	delete(data, key)

// 	return fs.Save(db.sourceName, data)
// }

// func (db *DB) Clear() error {
// 	return fs.Save(db.sourceName, map[string]string{})
// }

// func (db *DB) Update(key, value string) error {
// 	data, err := fs.Load(db.sourceName)

// 	if err != nil {
// 		return err
// 	}

// 	data[key] = value

// 	return fs.Save(db.sourceName, data)
// }
