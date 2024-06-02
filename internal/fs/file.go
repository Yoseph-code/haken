package fs

import "os"

func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsDirExist(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateDir(dirname string) error {
	if IsDirExist(dirname) {
		return nil
	}
	return os.Mkdir(dirname, 0755)
}

func CreateFile(filename string) error {
	if IsFileExist(filename) {
		return nil
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	return file.Close()
}
