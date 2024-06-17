package db

import (
	"bytes"
	"fmt"
)

type Data struct {
	Key   []byte
	Value []byte
}

func NewData(key, value []byte) *Data {
	return &Data{
		Key:   key,
		Value: value,
	}
}

func (data *Data) String() string {
	return fmt.Sprintf("%s=%s", data.Key, data.Value)
}

func (data *Data) Byte() []byte {
	return []byte(data.String())
}

type Node struct {
	Data  Data
	Left  *Node
	Right *Node
}

func NewNode(data Data) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) Insert(data Data) *Node {
	if bytes.Compare(data.Key, node.Data.Key) < 0 {
		if node.Left == nil {
			node.Left = NewNode(data)
		} else {
			node.Left.Insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(data)
		} else {
			node.Right.Insert(data)
		}
	}
	return node
}

func (node *Node) Append(data Data) {
	if bytes.Compare(data.Key, node.Data.Key) < 0 {
		if node.Left == nil {
			node.Left = NewNode(data)
		} else {
			node.Left.Append(data)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(data)
		} else {
			node.Right.Append(data)
		}
	}
}

func (node *Node) InOrderTraversal() {
	if node != nil {
		node.Left.InOrderTraversal()
		fmt.Println(node.Data.String())
		node.Right.InOrderTraversal()
	}
}

func (node *Node) Search(key []byte) *Node {
	if bytes.Equal(key, node.Data.Key) {
		return node
	} else if bytes.Compare(key, node.Data.Key) < 0 {
		if node.Left == nil {
			return nil
		}
		return node.Left.Search(key)
	} else {
		if node.Right == nil {
			return nil
		}
		return node.Right.Search(key)
	}
}

func (node *Node) Three() []byte {
	var data []byte
	if node != nil {
		data = append(data, node.Left.Three()...)
		data = append(data, node.Data.Byte()...)
		data = append(data, node.Right.Three()...)
	}
	return data
}
