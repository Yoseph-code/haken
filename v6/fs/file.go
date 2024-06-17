package fs

import "os"

const defaultFileMode = 0644

type File struct {
	Name string
	Data []byte
	Mode os.FileMode
}

func NewFile(name string, data []byte) *File {
	return &File{
		Name: name,
		Data: data,
		Mode: defaultFileMode,
	}
}

func (file *File) SetData(data []byte) {
	file.Data = data
}

func (file *File) Create() error {
	f, err := os.Create(file.Name)
	if err != nil {
		return err
	}

	return f.Close()
}

func (file *File) Exists() bool {
	_, err := os.Stat(file.Name)
	return err == nil
}

func (file *File) SetMode(mode os.FileMode) {
	file.Mode = mode
}

func (file *File) Write() error {
	f, err := os.Create(file.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(file.Data); err != nil {
		return err
	}

	return f.Chmod(file.Mode)
}

func (file *File) WriteAppend() error {
	f, err := os.OpenFile(file.Name, os.O_APPEND|os.O_WRONLY, file.Mode)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(file.Data); err != nil {
		return err
	}

	return nil
}

func (file *File) Read() error {
	f, err := os.Open(file.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	file.Data = make([]byte, stat.Size())
	_, err = f.Read(file.Data)
	if err != nil {
		return err
	}

	return nil
}

func (file *File) Remove() error {
	return os.Remove(file.Name)
}
