package fs

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"os"
)

func LoadFile(name string) ([]byte, int, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)

	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}

		if n == 0 {
			break
		}
	}

	return buf.Bytes(), len(buf.Bytes()), nil
}

func LoadFromFile(filename string, data interface{}) error {
	registerTypes()

	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(data)

	if err != nil {
		return err
	}

	return nil
}

func registerTypes() {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
	gob.Register("")   // string
	gob.Register(0)    // int
	gob.Register(0.0)  // float64
	gob.Register(true) // bool
}

func SaveToFile(filename string, data map[string]interface{}) error {
	registerTypes()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func AppendFile(filename string, data interface{}) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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
