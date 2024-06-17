package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

const (
	degree      = 4
	defaultFile = "data.bin"
	defaultMode = 0644
	keySize     = 32
	valueSize   = 256
	recordSize  = keySize + valueSize
	headerSize  = 4
	nodeSize    = headerSize + (degree-1)*recordSize + degree*4
)

type KeyValue struct {
	Key   []byte
	Value []byte
}

type Node struct {
	IsLeaf   bool
	NumKeys  int
	Keys     [][]byte
	Children []*Node
}

type BTree struct {
	Root *Node
}

func NewBTree() *BTree {
	return &BTree{
		Root: &Node{
			IsLeaf:   true,
			NumKeys:  0,
			Keys:     make([][]byte, degree-1),
			Children: make([]*Node, degree),
		},
	}
}

func (b *BTree) Insert(key, value []byte) {
	root := b.Root
	if root.NumKeys == 2*(degree-1) {
		newRoot := &Node{
			IsLeaf:   false,
			NumKeys:  0,
			Keys:     make([][]byte, degree-1),
			Children: make([]*Node, degree),
		}
		b.Root = newRoot
		newRoot.Children[0] = root
		b.splitChild(newRoot, 0)
		b.insertNonFull(newRoot, key, value)
	} else {
		b.insertNonFull(root, key, value)
	}
}

func (b *BTree) insertNonFull(node *Node, key, value []byte) {
	i := node.NumKeys - 1
	if node.IsLeaf {
		for i >= 0 && bytes.Compare(key, node.Keys[i]) < 0 {
			node.Keys[i+1] = node.Keys[i]
			i--
		}
		node.Keys[i+1] = key
		node.NumKeys++
	} else {
		for i >= 0 && bytes.Compare(key, node.Keys[i]) < 0 {
			i--
		}
		i++
		if node.Children[i].NumKeys == 2*(degree-1) {
			b.splitChild(node, i)
			if bytes.Compare(key, node.Keys[i]) > 0 {
				i++
			}
		}
		b.insertNonFull(node.Children[i], key, value)
	}
}

func (b *BTree) splitChild(node *Node, index int) {
	child := node.Children[index]
	newChild := &Node{
		IsLeaf:   child.IsLeaf,
		NumKeys:  degree - 1,
		Keys:     make([][]byte, degree-1),
		Children: make([]*Node, degree),
	}
	copy(newChild.Keys, child.Keys[degree:])
	if !child.IsLeaf {
		copy(newChild.Children, child.Children[degree:])
	}
	child.NumKeys = degree - 1
	for i := node.NumKeys; i > index; i-- {
		node.Children[i+1] = node.Children[i]
	}
	node.Children[index+1] = newChild
	for i := node.NumKeys - 1; i >= index; i-- {
		node.Keys[i+1] = node.Keys[i]
	}
	node.Keys[index] = child.Keys[degree-1]
	node.NumKeys++
}

func (b *BTree) Search(key []byte) ([]byte, error) {
	return b.search(b.Root, key)
}

func (b *BTree) search(node *Node, key []byte) ([]byte, error) {
	i := 0
	for i < node.NumKeys && bytes.Compare(key, node.Keys[i]) > 0 {
		i++
	}
	if i < node.NumKeys && bytes.Equal(key, node.Keys[i]) {
		return node.Keys[i], nil
	} else if node.IsLeaf {
		return nil, fmt.Errorf("key not found")
	} else {
		return b.search(node.Children[i], key)
	}
}

func (b *BTree) Traverse() {
	b.traverse(b.Root)
}

func (b *BTree) traverse(node *Node) {
	if node != nil {
		for i := 0; i < node.NumKeys; i++ {
			b.traverse(node.Children[i])
			fmt.Println(string(node.Keys[i]))
		}
		b.traverse(node.Children[node.NumKeys])
	}
}

func (b *BTree) WriteToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(b.Root)
	if err != nil {
		return err
	}

	return nil
}

func (b *BTree) ReadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&b.Root)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	btree := NewBTree()

	btree.Insert([]byte("D"), []byte("4"))
	btree.Insert([]byte("B"), []byte("2"))
	btree.Insert([]byte("A"), []byte("1"))
	btree.Insert([]byte("C"), []byte("3"))
	btree.Insert([]byte("F"), []byte("6"))
	btree.Insert([]byte("E"), []byte("5"))
	btree.Insert([]byte("G"), []byte("7"))
	btree.Insert([]byte("F"), []byte("8"))

	btree.Traverse()

	err := btree.WriteToFile(defaultFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("B-tree written to file successfully")

	err = btree.ReadFromFile(defaultFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := btree.Search([]byte("A"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Value:", string(value))
}
