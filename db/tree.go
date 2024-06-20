package db

type Node struct {
	Key   string
	Value []byte
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node

	fs *DBFile
}

func NewBinaryTree(node *Node, fs *DBFile) *BinaryTree {
	return &BinaryTree{
		Root: node,
		fs:   fs,
	}
}

func (bt *BinaryTree) Insert(key string, value []byte) {
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
}

func (bt *BinaryTree) insertNode(node *Node, key string, value []byte) {
	if key < node.Key {
		if node.Left == nil {
			node.Left = &Node{
				Key:   key,
				Value: value,
				Left:  nil,
				Right: nil,
			}
		} else {
			bt.insertNode(node.Left, key, value)
		}
	} else if key > node.Key {
		if node.Right == nil {
			node.Right = &Node{
				Key:   key,
				Value: value,
				Left:  nil,
				Right: nil,
			}
		} else {
			bt.insertNode(node.Right, key, value)
		}
	} else {
		node.Value = value
	}
}

func (bt *BinaryTree) Search(key string) (string, bool) {
	return bt.searchNode(bt.Root, key)
}

func (bt *BinaryTree) searchNode(node *Node, key string) (string, bool) {
	if node == nil {
		return "", false
	}

	if key < node.Key {
		return bt.searchNode(node.Left, key)
	} else if key > node.Key {
		return bt.searchNode(node.Right, key)
	} else {
		return string(node.Value), true
	}
}

func (bt *BinaryTree) Delete(key string) {
	bt.Root = bt.deleteNode(bt.Root, key)
}

func (bt *BinaryTree) deleteNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if key < node.Key {
		node.Left = bt.deleteNode(node.Left, key)
	} else if key > node.Key {
		node.Right = bt.deleteNode(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		node.Key = bt.minValue(node.Right)
		node.Right = bt.deleteNode(node.Right, node.Key)
	}

	return node
}

func (bt *BinaryTree) minValue(node *Node) string {
	minValue := node.Key
	for node.Left != nil {
		minValue = node.Left.Key
		node = node.Left
	}

	return minValue
}

func (bt *BinaryTree) Update(key string, value []byte) {
	bt.updateNode(bt.Root, key, value)
}

func (bt *BinaryTree) updateNode(node *Node, key string, value []byte) {
	if key < node.Key {
		bt.updateNode(node.Left, key, value)
	} else if key > node.Key {
		bt.updateNode(node.Right, key, value)
	} else {
		node.Value = value
	}
}

func (bt *BinaryTree) Print() {
	bt.printNode(bt.Root)
}

func (bt *BinaryTree) printNode(node *Node) {
	if node == nil {
		return
	}

	bt.printNode(node.Left)
	println(node.Key, node.Value)
	bt.printNode(node.Right)
}

func (bt *BinaryTree) InsertToFile(key string, value []byte) error {
	return bt.fs.Append(bt, key, value)
}

func (bt *BinaryTree) LoadFromFile(key string) (string, bool) {
	err := bt.fs.Load(bt, key)

	if err != nil {
		return "", false
	}

	return bt.Search(key)
}

func (bt *BinaryTree) UpdateToFile(key string, value []byte) error {
	return bt.fs.Update(bt, key, value)
}

func (bt *BinaryTree) RemoveFromFile(key string) error {
	return bt.fs.Remove(bt, key)
}
