package btree

type Node struct {
	Key   [32]byte
	Value [64]byte
	Left  *Node
	Right *Node
}

func NewNode(key [32]byte, value [64]byte) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}
