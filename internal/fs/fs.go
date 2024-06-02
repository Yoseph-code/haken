package fs

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func setKeyVal(k, v string) []byte {
	return []byte(fmt.Sprintf("%s=%s\n", k, v))
}

func Append(filename string, data map[string]string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer file.Close()

	old, err := Load(filename)

	if err != nil {
		return err
	}

	writer := gzip.NewWriter(file)

	defer writer.Close()

	for k, v := range data {
		if _, ok := old[k]; ok {
			return fmt.Errorf("key already exists")
		}

		if _, err = writer.Write(setKeyVal(k, v)); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	return writer.Flush()
}

func Load(filename string) (map[string]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data := make(map[string]string)

	reader, err := gzip.NewReader(file)

	if err != nil {
		if err == io.EOF {
			return data, nil
		}

		return nil, err
	}

	defer reader.Close()

	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)

		if err == io.EOF {
			break
		}

		if n == 0 {
			break
		}

		if err != nil {
			return nil, err
		}

		parts := strings.SplitN(string(buf[:n]), "=", 2)

		if len(parts) == 2 {
			data[parts[0]] = strings.TrimSuffix(parts[1], "\n")
		}
	}

	return data, nil
}
