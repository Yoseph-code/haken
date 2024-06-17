package fs

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	safeMap "github.com/Yoseph-code/haken/internal/safe_map"
)

const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

func init() {
	// node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	// assert(node1max <= BTREE_PAGE_SIZE) // maximum KV
}

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(data)

	if err != nil {
		return err
	}

	return fp.Sync()
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Intn(1000))

	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer func() {
		fp.Close()

		if err != nil {
			os.Remove(tmp)
		}
	}()

	_, err = fp.Write(data)

	if err != nil {
		return err
	}

	if err = fp.Sync(); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}

func setKeyVal(k, v string) []byte {
	return []byte(fmt.Sprintf("%s=%s\n", k, v))
}

func Append(filename string, key, value string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer file.Close()

	old, err := Load(filename)

	if err != nil {
		return err
	}

	defer old.Clear()

	writer := gzip.NewWriter(file)

	defer writer.Close()

	if ok := old.Has(key); ok {
		return fmt.Errorf("key already exists")
	}

	if _, err = writer.Write(setKeyVal(key, value)); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return writer.Flush()
}

func Load(filename string) (*safeMap.SafeMap[string], error) {
	// file, err := os.Open(filename)

	// if err != nil {
	// 	return nil, err
	// }

	// defer file.Close()

	// data := safeMap.NewSafeMap[string]()

	// reader, err := gzip.NewReader(file)

	// if err != nil {
	// 	if err == io.EOF {
	// 		return data, nil
	// 	}

	// 	return nil, err
	// }

	// defer reader.Close()

	// fi, err := file.Stat()

	// if err != nil {
	// 	return nil, err
	// }

	// buf := make([]byte, fi.Size())

	// for {
	// 	n, err := reader.Read(buf)

	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if n == 0 {
	// 		break
	// 	}

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	parts := bytes.SplitN(buf[:n], []byte{'='}, 2)

	// 	if len(parts) == 2 {
	// 		if ok := data.Has(string(parts[0])); !ok {
	// 			data.Set(string(parts[0]), string(bytes.TrimSuffix(parts[1], []byte{'\n'})))
	// 		}
	// 	}

	// 	copy(buf, buf[n:])
	// }

	// return data, nil

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	fmt.Println(file, err)

	return safeMap.NewSafeMap[string](), nil
}

// func preProcessBadChar(pattern []byte) [256]int {
// 	badChar := [256]int{}

// 	for i := range badChar {
// 		badChar[i] = -1
// 	}

// 	for i := 0; i < len(pattern); i++ {
// 		badChar[pattern[i]] = i
// 	}

// 	return badChar
// }

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func FindIndexOf(filename, pattern string) (int, error) {
// 	file, err := os.Open(filename)

// 	if err != nil {
// 		return -1, err
// 	}

// 	defer file.Close()

// 	reader, err := gzip.NewReader(file)

// 	if err != nil {
// 		return -1, err
// 	}

// 	defer reader.Close()

// 	fi, err := file.Stat()

// 	if err != nil {
// 		return -1, err
// 	}

// 	buf := make([]byte, fi.Size())

// 	for {
// 		n, err := reader.Read(buf)

// 		if err == io.EOF {
// 			break
// 		}

// 		if n == 0 {
// 			break
// 		}

// 		if err != nil {
// 			return -1, err
// 		}
// 	}

// 	fmt.Println(string(buf))

// 	return -1, nil

// 	// content, err := os.ReadFile(filename)

// 	// if err != nil {
// 	// 	return -1, err
// 	// }

// 	// badChar := preProcessBadChar([]byte(pattern))

// 	// s := 0

// 	// for s <= (len(content) - len(pattern)) {
// 	// 	j := len(pattern) - 1

// 	// 	for j >= 0 && pattern[j] == content[s+j] {
// 	// 		j--
// 	// 	}

// 	// 	if j < 0 {
// 	// 		return s, nil
// 	// 	} else {
// 	// 		s += Max(1, j-badChar[content[s+j]])
// 	// 	}
// 	// }

// 	// return -1, nil

// 	// // // //

// 	// file, err := os.Open(filename)

// 	// if err != nil {
// 	// 	return -1, err
// 	// }

// 	// defer file.Close()

// 	// reader, err := gzip.NewReader(file)

// 	// if err != nil {
// 	// 	return -1, err
// 	// }

// 	// defer reader.Close()

// 	// fi, err := file.Stat()

// 	// if err != nil {
// 	// 	return -1, err
// 	// }

// 	// buf := make([]byte, fi.Size())

// 	// for {
// 	// 	n, err := reader.Read(buf)

// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}

// 	// 	if n == 0 {
// 	// 		break
// 	// 	}

// 	// 	if err != nil {
// 	// 		return -1, err
// 	// 	}

// 	// 	parts := bytes.SplitN(buf[:n], []byte{'='}, 2)

// 	// 	if len(parts) == 2 {
// 	// 		if string(parts[0]) == key {
// 	// 			return 0, nil
// 	// 		}
// 	// 	}

// 	// 	copy(buf, buf[n:])
// 	// }

// 	// return -1, nil
// }

// Função para descomprimir um arquivo gzip em memória
func decompressGzipFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	return io.ReadAll(gzipReader)
}

// Função para preprocessar a tabela de bad character para o algoritmo Boyer-Moore
func preprocessBadChar(pattern []byte) [256]int {
	badChar := [256]int{}
	for i := range badChar {
		badChar[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		badChar[pattern[i]] = i
	}
	return badChar
}

// Implementa a busca Boyer-Moore
func BoyerMooreSearch(content, pattern []byte) int {
	badChar := preprocessBadChar(pattern)
	s := 0
	for s <= (len(content) - len(pattern)) {
		j := len(pattern) - 1

		for j >= 0 && pattern[j] == content[s+j] {
			j--
		}
		if j < 0 {
			return s
		} else {
			s += max(1, j-badChar[content[s+j]])
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Exec() {
	// Descomprimir o arquivo gzip
	content, err := decompressGzipFile("/Users/eco/Developer/haken/haken/default.db")
	if err != nil {
		log.Fatal(err)
	}

	// Padrão a ser buscado
	pattern := []byte("json1")

	// Executar a busca Boyer-Moore
	pos := BoyerMooreSearch(content, pattern)
	if pos != -1 {
		fmt.Printf("Padrão encontrado na posição %d\n", pos)
	} else {
		fmt.Println("Padrão não encontrado")
	}
}
