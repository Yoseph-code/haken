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
	if _, err := f.Read(file.Data); err != nil {
		return err
	}

	return nil
}

func (file *File) Remove() error {
	return os.Remove(file.Name)
}

func (file *File) Exists() bool {
	_, err := os.Stat(file.Name)
	return !os.IsNotExist(err)
}

func (file *File) Rename(newName string) error {
	return os.Rename(file.Name, newName)
}

func (file *File) Copy(newName string) error {
	f, err := os.Create(newName)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(file.Data); err != nil {
		return err
	}

	return f.Chmod(file.Mode)
}

func (file *File) Move(newName string) error {
	if err := file.Copy(newName); err != nil {
		return err
	}

	return file.Remove()
}

func (file *File) Size() int64 {
	return int64(len(file.Data))
}

func (file *File) IsDir() bool {
	stat, err := os.Stat(file.Name)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

func (file *File) IsFile() bool {
	stat, err := os.Stat(file.Name)
	if err != nil {
		return false
	}

	return stat.Mode().IsRegular()
}

func (file *File) IsSymlink() bool {
	stat, err := os.Lstat(file.Name)
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeSymlink != 0
}

func (file *File) IsSocket() bool {
	stat, err := os.Stat(file.Name)
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeSocket != 0
}

func (file *File) IsNamedPipe() bool {
	stat, err := os.Stat(file.Name)
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeNamedPipe != 0
}

func (file *File) IsCharDevice() bool {
	stat, err := os.Stat(file.Name)
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeCharDevice != 0
}

func (file *File) IsBlockDevice() bool {
	stat, err := os.Stat(file.Name)
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeDevice != 0
}

func (file *File) IsReadable() bool {
	return file.Mode&0400 != 0
}

func (file *File) IsWritable() bool {
	return file.Mode&0200 != 0
}

func (file *File) IsExecutable() bool {
	return file.Mode&0100 != 0
}

func (file *File) IsSetuid() bool {
	return file.Mode&04000 != 0
}

func (file *File) IsSetgid() bool {
	return file.Mode&02000 != 0
}

func (file *File) IsSticky() bool {
	return file.Mode&01000 != 0
}

func (file *File) IsRegular() bool {
	return file.Mode&os.ModeType == 0
}

func (file *File) Create() error {
	f, err := os.Create(file.Name)

	if err != nil {
		return err
	}

	return f.Close()
}
