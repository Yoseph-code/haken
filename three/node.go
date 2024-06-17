package three

type Node struct {
	isLeaf   bool
	keys     []int
	children []*Node
	values   []interface{}
}

type BPlusTree struct {
	root *Node
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		root: &Node{
			isLeaf:   true,
			keys:     []int{},
			children: []*Node{},
			values:   []interface{}{},
		},
	}
}

func (t *BPlusTree) Insert(key int, value interface{}) {
	leaf := t.findLeafNode(key)
	leaf.insertKey(key, value)
}

func (t *BPlusTree) Find(key int) (interface{}, bool) {
	leaf := t.findLeafNode(key)
	return leaf.findValue(key)
}

func (t *BPlusTree) Delete(key int) bool {
	leaf := t.findLeafNode(key)
	return leaf.deleteKey(key)
}

func (t *BPlusTree) findLeafNode(key int) *Node {
	node := t.root
	for !node.isLeaf {
		i := 0
		for i < len(node.keys) && key >= node.keys[i] {
			i++
		}
		node = node.children[i]
	}
	return node
}

func (n *Node) insertKey(key int, value interface{}) {
	i := 0
	for i < len(n.keys) && key > n.keys[i] {
		i++
	}
	n.keys = append(n.keys[:i], append([]int{key}, n.keys[i:]...)...)
	n.values = append(n.values[:i], append([]interface{}{value}, n.values[i:]...)...)
}

func (n *Node) findValue(key int) (interface{}, bool) {
	i := 0
	for i < len(n.keys) && key > n.keys[i] {
		i++
	}
	if i < len(n.keys) && key == n.keys[i] {
		return n.values[i], true
	}
	return nil, false
}

func (n *Node) deleteKey(key int) bool {
	i := 0
	for i < len(n.keys) && key > n.keys[i] {
		i++
	}
	if i < len(n.keys) && key == n.keys[i] {
		n.keys = append(n.keys[:i], n.keys[i+1:]...)
		n.values = append(n.values[:i], n.values[i+1:]...)
		return true
	}
	return false
}

func (t *BPlusTree) Update(key int, value interface{}) bool {
	leaf := t.findLeafNode(key)
	return leaf.updateValue(key, value)
}

func (n *Node) updateValue(key int, value interface{}) bool {
	i := 0
	for i < len(n.keys) && key > n.keys[i] {
		i++
	}
	if i < len(n.keys) && key == n.keys[i] {
		n.values[i] = value
		return true
	}
	return false
}

func (t *BPlusTree) Range(from, to int) []interface{} {
	result := []interface{}{}
	node := t.findLeafNode(from)
	for node != nil {
		for i := 0; i < len(node.keys); i++ {
			if node.keys[i] >= from && node.keys[i] <= to {
				result = append(result, node.values[i])
			}
		}
		node = t.findNextLeafNode(node)
	}
	return result
}

func (t *BPlusTree) findNextLeafNode(node *Node) *Node {
	if node.children == nil {
		return nil
	}
	return node.children[0]
}

func (t *BPlusTree) FindAll() []interface{} {
	result := []interface{}{}
	node := t.findFirstLeafNode()
	for node != nil {
		result = append(result, node.values...)
		node = t.findNextLeafNode(node)
	}
	return result
}

func (t *BPlusTree) findFirstLeafNode() *Node {
	node := t.root
	for !node.isLeaf {
		node = node.children[0]
	}
	return node
}

func (t *BPlusTree) FindRange(from, to int) []interface{} {
	result := []interface{}{}
	node := t.findLeafNode(from)
	for node != nil {
		for i := 0; i < len(node.keys); i++ {
			if node.keys[i] >= from && node.keys[i] <= to {
				result = append(result, node.values[i])
			}
		}
		node = t.findNextLeafNode(node)
	}
	return result
}

func (t *BPlusTree) FindAllRange(from, to int) []interface{} {
	result := []interface{}{}
	node := t.findFirstLeafNode()
	for node != nil {
		for i := 0; i < len(node.keys); i++ {
			if node.keys[i] >= from && node.keys[i] <= to {
				result = append(result, node.values[i])
			}
		}
		node = t.findNextLeafNode(node)
	}
	return result
}
