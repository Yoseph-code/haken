package fs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	safeMap "github.com/Yoseph-code/haken/internal/safe_map"
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

	// old, err := Load(filename)

	// if err != nil {
	// 	return err
	// }

	writer := gzip.NewWriter(file)

	defer writer.Close()

	// for k, v := range data {
	// 	if _, ok := old[k]; ok {
	// 		return fmt.Errorf("key already exists")
	// 	}

	// 	if _, err = writer.Write(setKeyVal(k, v)); err != nil {
	// 		return fmt.Errorf("failed to write to file: %v", err)
	// 	}
	// }

	return writer.Flush()
}

func Load(filename string) (*safeMap.SafeMap[string], error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data := safeMap.NewSafeMap[string]()

	reader, err := gzip.NewReader(file)

	if err != nil {
		if err == io.EOF {
			return data, nil
		}

		return nil, err
	}

	defer reader.Close()

	fi, err := file.Stat()

	if err != nil {
		return nil, err
	}

	buf := make([]byte, fi.Size())

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

		parts := bytes.SplitN(buf[:n], []byte{'='}, 2)

		if len(parts) == 2 {
			if ok := data.Has(string(parts[0])); !ok {
				data.Set(string(parts[0]), string(bytes.TrimSuffix(parts[1], []byte{'\n'})))
			}
		}

		copy(buf, buf[n:])
	}

	return data, nil
}
