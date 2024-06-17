package db

import (
	"encoding/gob"
	"os"

	"github.com/Yoseph-code/haken/three"
)

type DB struct {
	filename string

	Data *three.BThree
}

func NewDB(filename string, t int) *DB {
	return &DB{
		filename: filename,
		Data:     three.NewBThree(t),
	}
}

func (db *DB) WriteToFile(data *three.Node) error {
	file, err := os.OpenFile(db.filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := gob.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) AppendToFile(data *three.Node) error {
	file, err := os.OpenFile(db.filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := gob.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) ReadFromFile() error {
	file, err := os.Open(db.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(&db.Data.Root)

	if err != nil {
		return err
	}

	return nil
}
