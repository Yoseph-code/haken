package db

import (
	"errors"

	"github.com/Yoseph-code/haken/internal/fs"
)

func (db *DB) Create(key, value string) error {
	data, err := fs.Load(db.sourceName)

	if err != nil {
		return err
	}

	defer data.Clear()

	if ok := data.Has(key); ok {
		return errors.New("key already exists")
	}

	return fs.Append(db.sourceName, key, value)
}

// TODO: implement this
// func (db *DB) ReadAll() (map[string]string, error) {
// 	data, err := fs.Load(db.sourceName)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.Copy(), nil
// }

func (db *DB) Read(key string) (string, error) {
	data, err := fs.Load(db.sourceName)

	if err != nil {
		return "", err
	}

	defer data.Clear()

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
