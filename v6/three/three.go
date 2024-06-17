package three

import "fmt"

type BThree struct {
	Root *Node
	t    int
}

func NewBThree(t int) *BThree {
	return &BThree{
		Root: NewNode(t, true),
		t:    t,
	}
}

func (bt *BThree) Insert(key, value string) {
	root := bt.Root

	if len(root.Keys) == 2*bt.t-1 {
		newRoot := NewNode(bt.t, false)
		newRoot.Children = append(newRoot.Children, root)
		bt.splitChild(newRoot, 0, root)
		bt.Root = newRoot
		bt.insertNonFull(newRoot, key, value)
	} else {
		bt.insertNonFull(root, key, value)
	}
}

func (bt *BThree) insertNonFull(node *Node, key, value string) {
	i := len(node.Keys) - 1

	if node.IsLeaf {
		node.Keys = append(node.Keys, "")
		node.Values = append(node.Values, "")

		for i >= 0 && key < node.Keys[i] {
			node.Keys[i+1] = node.Keys[i]
			node.Values[i+1] = node.Values[i]

			i--
		}

		node.Keys[i+1] = key
		node.Values[i+1] = value
	} else {
		for i >= 0 && key < node.Keys[i] {
			i--
		}

		i++

		if len(node.Children[i].Keys) == 2*bt.t-1 {
			bt.splitChild(node, i, node.Children[i])

			if key > node.Keys[i] {
				i++
			}
		}

		bt.insertNonFull(node.Children[i], key, value)
	}
}

func (bt *BThree) splitChild(parent *Node, i int, child *Node) {
	t := bt.t

	newChild := NewNode(t, child.IsLeaf)

	parent.Children = append(parent.Children[:i+1], append([]*Node{newChild}, parent.Children[i+1:]...)...)

	parent.Keys = append(parent.Keys[:i], append([]string{child.Keys[t-1]}, parent.Keys[i:]...)...)

	parent.Values = append(parent.Values[:i], append([]string{child.Values[t-1]}, parent.Values[i:]...)...)

	newChild.Keys = append(newChild.Keys, child.Keys[t:]...)

	newChild.Values = append(newChild.Values, child.Values[t:]...)

	child.Keys = child.Keys[:t-1]

	child.Values = child.Values[:t-1]

	if !child.IsLeaf {
		newChild.Children = append(newChild.Children, child.Children[t:]...)
		child.Children = child.Children[:t]
	}

	if child.IsLeaf {
		newChild.Next = child.Next
		child.Next = newChild
	}
}

func (bt *BThree) Print() {
	printNode(bt.Root, 0)
}

func printNode(node *Node, level int) {
	fmt.Printf("%sLevel %d %v\n", string(rune(level*4)), level, node.Keys)

	if node.IsLeaf {
		fmt.Printf("%sValues %v\n", string(rune(level*4)), node.Values)
	}

	for _, child := range node.Children {
		printNode(child, level+1)
	}
}

func (bt *BThree) Search(key string) (string, bool) {
	return bt.search(bt.Root, key)
}

func (bt *BThree) search(node *Node, key string) (string, bool) {
	i := 0

	for i < len(node.Keys) && key > node.Keys[i] {
		i++
	}

	if i < len(node.Keys) && key == node.Keys[i] {
		return node.Values[i], true
	}

	if node.IsLeaf {
		return "", false
	}

	return bt.search(node.Children[i], key)
}
