package btree

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree(node *Node) *BinaryTree {
	return &BinaryTree{
		Root: node,
	}
}
