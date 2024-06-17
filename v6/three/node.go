package three

type Node struct {
	IsLeaf   bool
	Keys     []string
	Values   []string
	Children []*Node
	Next     *Node
}

func NewNode(t int, isLeaf bool) *Node {
	return &Node{
		IsLeaf:   isLeaf,
		Keys:     make([]string, 0, 2*t-1),
		Values:   make([]string, 0, 2*t-1),
		Children: make([]*Node, 0, 2*t),
	}
}
