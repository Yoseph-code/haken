package db

type Three struct {
	Root *Node
	T    int
}

func NewThree(t int, isLeaf bool) *Three {
	return &Three{
		Root: NewNode(t, isLeaf),
		T:    t,
	}
}
