package db

import (
	"encoding/binary"
	"os"
	"path"
)

const (
	mainPath string = "haken"
	mainFile string = "haken.bin"
)

type DBFile struct {
	filename string
}

func NewDBFile() (*DBFile, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	filename := path.Join(pwd, mainPath, mainFile)

	return &DBFile{
		filename: filename,
	}, nil
}

func (db *DBFile) IsDBExists() bool {
	_, err := os.Stat(db.filename)

	return !os.IsNotExist(err)
}

func (db *DBFile) CreateDB() (*os.File, error) {
	if ok := db.IsDBExists(); !ok {
		return os.OpenFile(db.filename, os.O_CREATE, 0644)
	}

	return os.Create(db.filename)
}

func (db *DBFile) Way() string {
	return db.filename
}

func (fs *DBFile) Append(bt *BinaryTree, key string, value []byte) error {
	file, err := os.OpenFile(fs.filename, os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	keyBytes := []byte(key)

	valueLen := make([]byte, 4)
	binary.BigEndian.PutUint32(valueLen, uint32(len(value)))

	if err := binary.Write(file, binary.BigEndian, uint32(len(keyBytes))); err != nil {
		return err
	}
	if _, err := file.Write(keyBytes); err != nil {
		return err
	}
	if err := binary.Write(file, binary.BigEndian, valueLen); err != nil {
		return err
	}
	if _, err := file.Write(value); err != nil {
		return err
	}

	if bt.Root == nil {
		bt.Root = &Node{
			Key:   key,
			Value: value,
			Left:  nil,
			Right: nil,
		}
	} else {
		bt.insertNode(bt.Root, key, value)
	}

	return nil
}

func (fs *DBFile) Load(bt *BinaryTree) error {
	file, err := os.Open(fs.filename)

	if err != nil {
		return err
	}

	defer file.Close()

	var keyLen uint32

	for {
		err := binary.Read(file, binary.BigEndian, &keyLen)

		if err != nil {
			break
		}

		key := make([]byte, keyLen)

		_, err = file.Read(key)

		if err != nil {
			break
		}

		var valueLen uint32

		err = binary.Read(file, binary.BigEndian, &valueLen)

		if err != nil {
			break
		}

		value := make([]byte, valueLen)

		_, err = file.Read(value)

		if err != nil {
			break
		}

		if bt.Root == nil {
			bt.Root = &Node{
				Key:   string(key),
				Value: value,
				Left:  nil,
				Right: nil,
			}
		} else {
			bt.insertNode(bt.Root, string(key), value)
		}
	}

	return nil
}

// func (fs *DBFile) Load(bt *BinaryTree, props string) error {
// 	file, err := os.Open(fs.filename)

// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	var keyLen uint32

// 	for {
// 		err := binary.Read(file, binary.BigEndian, &keyLen)

// 		if err != nil {
// 			break
// 		}

// 		key := make([]byte, keyLen)

// 		_, err = file.Read(key)

// 		if err != nil {
// 			break
// 		}

// 		if string(key) == props {
// 			var valueLen uint32

// 			err = binary.Read(file, binary.BigEndian, &valueLen)

// 			if err != nil {
// 				break
// 			}

// 			value := make([]byte, valueLen)

// 			_, err = file.Read(value)

// 			if err != nil {
// 				break
// 			}

// 			if bt.Root == nil {
// 				bt.Root = &Node{
// 					Key:   string(key),
// 					Value: value,
// 					Left:  nil,
// 					Right: nil,
// 				}

// 				break

// 			} else {
// 				bt.insertNode(bt.Root, string(key), value)

// 				break
// 			}
// 		}

// 		var valueLen uint32

// 		err = binary.Read(file, binary.BigEndian, &valueLen)

// 		if err != nil {

// 			break
// 		}

// 		value := make([]byte, valueLen)

// 		_, err = file.Read(value)

// 		if err != nil {
// 			break
// 		}

// 		if bt.Root == nil {
// 			bt.Root = &Node{
// 				Key:   string(key),
// 				Value: value,
// 				Left:  nil,
// 				Right: nil,
// 			}

// 			break

// 		} else {
// 			bt.insertNode(bt.Root, string(key), value)

// 			break
// 		}
// 	}

// 	return nil
// }

// func (fs *DBFile) Update(bt *BinaryTree, targetKey string, value []byte) error {
// 	file, err := os.OpenFile(fs.filename, os.O_RDWR, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	var keyLen uint32
// 	var offset int64
// 	var found bool

// 	for {
// 		err := binary.Read(file, binary.BigEndian, &keyLen)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			return err
// 		}
// 		key := make([]byte, keyLen)
// 		_, err = file.Read(key)
// 		if err != nil {
// 			return err
// 		}

// 		if string(key) == targetKey {
// 			found = true
// 			var valueLen uint32
// 			err = binary.Read(file, binary.BigEndian, &valueLen)
// 			if err != nil {
// 				return err
// 			}
// 			oldValue := make([]byte, valueLen)
// 			_, err = file.Read(oldValue)
// 			if err != nil {
// 				return err
// 			}
// 			bt.updateNode(bt.Root, string(key), value)
// 			offset, _ = file.Seek(0, io.SeekCurrent)
// 			file.Seek(offset-int64(valueLen)-int64(keyLen)-8, io.SeekStart)
// 			binary.Write(file, binary.BigEndian, uint32(len(value)))
// 			file.Write(value)
// 			break
// 		} else {
// 			var valueLen uint32
// 			err = binary.Read(file, binary.BigEndian, &valueLen)
// 			if err != nil {
// 				return err
// 			}
// 			value := make([]byte, valueLen)
// 			_, err = file.Read(value)
// 			if err != nil {
// 				return err
// 			}
// 			if bt.Root == nil {
// 				bt.Root = &Node{
// 					Key:   string(key),
// 					Value: value,
// 					Left:  nil,
// 					Right: nil,
// 				}
// 			} else {
// 				bt.insertNode(bt.Root, string(key), value)
// 			}
// 		}
// 	}

// 	if !found {
// 		return fmt.Errorf("key not found: %s", targetKey)
// 	}

// 	return nil
// }

func (fs *DBFile) Update(bt *BinaryTree, targetKey string, value []byte) error {
	bt.Update(targetKey, value)

	return fs.Insert(bt)
}

func (fs *DBFile) Insert(bt *BinaryTree) error {
	file, err := os.OpenFile(fs.filename, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = fs.traverseAndAppend(bt.Root, file)
	if err != nil {
		return err
	}

	return nil
}

func (fs *DBFile) traverseAndAppend(node *Node, file *os.File) error {
	if node == nil {
		return nil
	}

	keyBytes := []byte(node.Key)
	valueLen := make([]byte, 4)
	binary.BigEndian.PutUint32(valueLen, uint32(len(node.Value)))

	if err := binary.Write(file, binary.BigEndian, uint32(len(keyBytes))); err != nil {
		return err
	}
	if _, err := file.Write(keyBytes); err != nil {
		return err
	}
	if err := binary.Write(file, binary.BigEndian, valueLen); err != nil {
		return err
	}
	if _, err := file.Write(node.Value); err != nil {
		return err
	}

	err := fs.traverseAndAppend(node.Left, file)
	if err != nil {
		return err
	}
	err = fs.traverseAndAppend(node.Right, file)
	if err != nil {
		return err
	}

	return nil
}

func (fs *DBFile) Remove(bt *BinaryTree, targetKey string) error {
	bt.Delete(targetKey)

	return fs.Insert(bt)
}
