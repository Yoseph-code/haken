package db

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
)

func (d *DB) Read() ([]byte, error) {
	file, err := os.OpenFile(*d.db, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	buf := new(bytes.Buffer)

	for {
		var size int64

		binary.Read(file, binary.LittleEndian, &size)

		n, err := io.CopyN(buf, file, size)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if n == 0 {
			break
		}
	}

	return buf.Bytes(), nil
}

func (d *DB) Write(data []byte) error {
	file, err := os.OpenFile(*d.db, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	size := int64(len(data))

	binary.Write(file, binary.LittleEndian, size)

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	return nil
}
