package hf

import (
	"io"
	"os"
	"path/filepath"
)

type HakenFile struct {
	filename string

	file *os.File
}

func NewHakenFile() *HakenFile {
	return &HakenFile{}
}

func (hf *HakenFile) SetFileName(filename string) error {
	pwd, err := os.Getwd()

	if err != nil {
		return err
	}

	hf.filename = filepath.Join(pwd, mainDir, filename)

	return nil
}

func (hf *HakenFile) CreateOrRead() (err error) {
	if _, err := os.Stat(hf.filename); os.IsNotExist(err) {
		hf.file, err = os.Create(hf.filename)

		if err != nil {
			return err
		}
	}

	hf.file, err = os.OpenFile(hf.filename, os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (hf *HakenFile) AppendToFile(data []byte) error {
	_, err := hf.file.Seek(0, io.SeekEnd)

	if err != nil {
		return err
	}

	_, err = hf.file.Write(data)

	return err
}

func (hf *HakenFile) Close() error {
	return hf.file.Close()
}

func (hf *HakenFile) ReadAll() ([]byte, error) {
	return os.ReadFile(hf.filename)
}
