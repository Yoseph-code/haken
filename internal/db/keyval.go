package db

import (
	"bytes"
	"encoding/gob"
	"errors"
	"strings"
)

func (db *DB) encodeToBinary(data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (db *DB) decodeFromBinary(binaryData []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	buf := bytes.NewBuffer(binaryData)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *DB) Get(key string) (interface{}, bool) {
	b, err := db.Read()

	if err != nil {
		return nil, false
	}

	data, err := db.decodeFromBinary(b)

	if err != nil {
		return nil, false
	}

	val, ok := data[key]

	return val, ok
}

func (db *DB) Set(key string, val []byte) error {
	b, err := db.Read()

	if err != nil {
		return err
	}

	data, err := db.decodeFromBinary(b)

	if err != nil {
		return err
	}

	if _, ok := data[key]; ok {
		return errors.New("key already exists")
	}

	data[key] = strings.TrimSpace(string(val))

	b, err = db.encodeToBinary(data)

	if err != nil {
		return err
	}

	return db.Write(b)
}
