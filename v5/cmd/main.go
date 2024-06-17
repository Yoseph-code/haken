// package main

// import "fmt"

// // Node representa um nó na árvore B+
// type Node struct {
// 	isLeaf   bool
// 	keys     []int
// 	children []*Node
// 	next     *Node // Usado para ligar folhas
// }

// // BPlusTree representa a árvore B+
// type BPlusTree struct {
// 	root *Node
// 	t    int // Grau mínimo
// }

// // Nova instância de nó
// func newNode(t int, isLeaf bool) *Node {
// 	return &Node{
// 		isLeaf:   isLeaf,
// 		keys:     make([]int, 0, 2*t-1),
// 		children: make([]*Node, 0, 2*t),
// 	}
// }

// // Nova instância de árvore B+
// func newBPlusTree(t int) *BPlusTree {
// 	root := newNode(t, true)
// 	return &BPlusTree{root: root, t: t}
// }

// // Insere uma chave na árvore B+
// func (tree *BPlusTree) insert(key int) {
// 	root := tree.root
// 	if len(root.keys) == 2*tree.t-1 {
// 		newRoot := newNode(tree.t, false)
// 		newRoot.children = append(newRoot.children, root)
// 		tree.splitChild(newRoot, 0, root)
// 		tree.root = newRoot
// 		tree.insertNonFull(newRoot, key)
// 	} else {
// 		tree.insertNonFull(root, key)
// 	}
// }

// // Insere uma chave em um nó que não está cheio
// func (tree *BPlusTree) insertNonFull(node *Node, key int) {
// 	i := len(node.keys) - 1
// 	if node.isLeaf {
// 		node.keys = append(node.keys, 0)
// 		for i >= 0 && key < node.keys[i] {
// 			node.keys[i+1] = node.keys[i]
// 			i--
// 		}
// 		node.keys[i+1] = key
// 	} else {
// 		for i >= 0 && key < node.keys[i] {
// 			i--
// 		}
// 		i++
// 		if len(node.children[i].keys) == 2*tree.t-1 {
// 			tree.splitChild(node, i, node.children[i])
// 			if key > node.keys[i] {
// 				i++
// 			}
// 		}
// 		tree.insertNonFull(node.children[i], key)
// 	}
// }

// // Divide um nó filho
// func (tree *BPlusTree) splitChild(parent *Node, i int, child *Node) {
// 	t := tree.t
// 	newChild := newNode(t, child.isLeaf)
// 	parent.children = append(parent.children[:i+1], append([]*Node{newChild}, parent.children[i+1:]...)...)
// 	parent.keys = append(parent.keys[:i], append([]int{child.keys[t-1]}, parent.keys[i:]...)...)
// 	newChild.keys = append(newChild.keys, child.keys[t:]...)
// 	child.keys = child.keys[:t-1]
// 	if !child.isLeaf {
// 		newChild.children = append(newChild.children, child.children[t:]...)
// 		child.children = child.children[:t]
// 	}
// 	if child.isLeaf {
// 		newChild.next = child.next
// 		child.next = newChild
// 	}
// }

// // Imprime a árvore B+
// func (tree *BPlusTree) print() {
// 	printNode(tree.root, 0)
// }

// // Função auxiliar para imprimir um nó
// func printNode(node *Node, level int) {
// 	fmt.Printf("%sLevel %d %v\n", string(level*4), level, node.keys)
// 	for _, child := range node.children {
// 		printNode(child, level+1)
// 	}
// }

// // Função principal para testar a árvore B+
// func main() {
// 	t := 3 // Grau mínimo
// 	bptree := newBPlusTree(t)

// 	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
// 	for _, key := range keys {
// 		bptree.insert(key)
// 	}

// 	fmt.Println("Árvore B+:")
// 	bptree.print()
// }

// package main

// import (
// 	"fmt"
// )

// // Node representa um nó na árvore B+
// type Node struct {
// 	isLeaf   bool
// 	keys     []interface{}
// 	children []*Node
// 	next     *Node // Usado para ligar folhas
// }

// // BPlusTree representa a árvore B+
// type BPlusTree struct {
// 	root *Node
// 	t    int // Grau mínimo
// }

// // Nova instância de nó
// func newNode(t int, isLeaf bool) *Node {
// 	return &Node{
// 		isLeaf:   isLeaf,
// 		keys:     make([]interface{}, 0, 2*t-1),
// 		children: make([]*Node, 0, 2*t),
// 	}
// }

// // Nova instância de árvore B+
// func newBPlusTree(t int) *BPlusTree {
// 	root := newNode(t, true)
// 	return &BPlusTree{root: root, t: t}
// }

// func compare(a, b interface{}) int {
// 	switch a.(type) {
// 	case int:
// 		ai := a.(int)
// 		bi := b.(int)
// 		if ai < bi {
// 			return -1
// 		} else if ai > bi {
// 			return 1
// 		}
// 	case float64:
// 		af := a.(float64)
// 		bf := b.(float64)
// 		if af < bf {
// 			return -1
// 		} else if af > bf {
// 			return 1
// 		}
// 	case string:
// 		as := a.(string)
// 		bs := b.(string)
// 		if as < bs {
// 			return -1
// 		} else if as > bs {
// 			return 1
// 		}
// 	}
// 	return 0
// }

// // Insere uma chave na árvore B+
// func (tree *BPlusTree) insert(key interface{}) {
// 	root := tree.root
// 	if len(root.keys) == 2*tree.t-1 {
// 		newRoot := newNode(tree.t, false)
// 		newRoot.children = append(newRoot.children, root)
// 		tree.splitChild(newRoot, 0, root)
// 		tree.root = newRoot
// 		tree.insertNonFull(newRoot, key)
// 	} else {
// 		tree.insertNonFull(root, key)
// 	}
// }

// // Insere uma chave em um nó que não está cheio
// func (tree *BPlusTree) insertNonFull(node *Node, key interface{}) {
// 	i := len(node.keys) - 1
// 	if node.isLeaf {
// 		node.keys = append(node.keys, nil)
// 		for i >= 0 && compare(key, node.keys[i]) < 0 {
// 			node.keys[i+1] = node.keys[i]
// 			i--
// 		}
// 		node.keys[i+1] = key
// 	} else {
// 		for i >= 0 && compare(key, node.keys[i]) < 0 {
// 			i--
// 		}
// 		i++
// 		if len(node.children[i].keys) == 2*tree.t-1 {
// 			tree.splitChild(node, i, node.children[i])
// 			if compare(key, node.keys[i]) > 0 {
// 				i++
// 			}
// 		}
// 		tree.insertNonFull(node.children[i], key)
// 	}
// }

// // Divide um nó filho
// func (tree *BPlusTree) splitChild(parent *Node, i int, child *Node) {
// 	t := tree.t
// 	newChild := newNode(t, child.isLeaf)
// 	parent.children = append(parent.children[:i+1], append([]*Node{newChild}, parent.children[i+1:]...)...)
// 	parent.keys = append(parent.keys[:i], append([]interface{}{child.keys[t-1]}, parent.keys[i:]...)...)
// 	newChild.keys = append(newChild.keys, child.keys[t:]...)
// 	child.keys = child.keys[:t-1]
// 	if !child.isLeaf {
// 		newChild.children = append(newChild.children, child.children[t:]...)
// 		child.children = child.children[:t]
// 	}
// 	if child.isLeaf {
// 		newChild.next = child.next
// 		child.next = newChild
// 	}
// }

// // Imprime a árvore B+
// func (tree *BPlusTree) print() {
// 	printNode(tree.root, 0)
// }

// // Função auxiliar para imprimir um nó
// func printNode(node *Node, level int) {
// 	fmt.Printf("%s Level %d %v\n", string(rune(level*4)), level, node.keys)
// 	for _, child := range node.children {
// 		printNode(child, level+1)
// 	}
// }

// // Função principal para testar a árvore B+
// func main() {
// 	t := 3 // Grau mínimo
// 	bptree := newBPlusTree(t)

// 	// Inserindo diferentes tipos de dados
// 	keys := []interface{}{10, 20.5, "five", 6, "twelve", 30.1, "seven", 17}
// 	for _, key := range keys {
// 		bptree.insert(key)
// 	}

// 	fmt.Println("Árvore B+:")
// 	bptree.print()
// }

package main

import (
	"fmt"
)

// Node representa um nó na árvore B+
type Node struct {
	isLeaf   bool
	keys     []string
	values   []string
	children []*Node
	next     *Node // Usado para ligar folhas
}

// BPlusTree representa a árvore B+
type BPlusTree struct {
	root *Node
	t    int // Grau mínimo
}

// Nova instância de nó
func newNode(t int, isLeaf bool) *Node {
	return &Node{
		isLeaf:   isLeaf,
		keys:     make([]string, 0, 2*t-1),
		values:   make([]string, 0, 2*t-1),
		children: make([]*Node, 0, 2*t),
	}
}

// Nova instância de árvore B+
func newBPlusTree(t int) *BPlusTree {
	root := newNode(t, true)
	return &BPlusTree{root: root, t: t}
}

// Insere um par chave-valor na árvore B+
func (tree *BPlusTree) insert(key, value string) {
	root := tree.root
	if len(root.keys) == 2*tree.t-1 {
		newRoot := newNode(tree.t, false)
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0, root)
		tree.root = newRoot
		tree.insertNonFull(newRoot, key, value)
	} else {
		tree.insertNonFull(root, key, value)
	}
}

// Insere um par chave-valor em um nó que não está cheio
func (tree *BPlusTree) insertNonFull(node *Node, key, value string) {
	i := len(node.keys) - 1
	if node.isLeaf {
		node.keys = append(node.keys, "")
		node.values = append(node.values, "")
		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			node.values[i+1] = node.values[i]
			i--
		}
		node.keys[i+1] = key
		node.values[i+1] = value
	} else {
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++
		if len(node.children[i].keys) == 2*tree.t-1 {
			tree.splitChild(node, i, node.children[i])
			if key > node.keys[i] {
				i++
			}
		}
		tree.insertNonFull(node.children[i], key, value)
	}
}

// Divide um nó filho
func (tree *BPlusTree) splitChild(parent *Node, i int, child *Node) {
	t := tree.t
	newChild := newNode(t, child.isLeaf)
	parent.children = append(parent.children[:i+1], append([]*Node{newChild}, parent.children[i+1:]...)...)
	parent.keys = append(parent.keys[:i], append([]string{child.keys[t-1]}, parent.keys[i:]...)...)
	parent.values = append(parent.values[:i], append([]string{child.values[t-1]}, parent.values[i:]...)...)
	newChild.keys = append(newChild.keys, child.keys[t:]...)
	newChild.values = append(newChild.values, child.values[t:]...)
	child.keys = child.keys[:t-1]
	child.values = child.values[:t-1]
	if !child.isLeaf {
		newChild.children = append(newChild.children, child.children[t:]...)
		child.children = child.children[:t]
	}
	if child.isLeaf {
		newChild.next = child.next
		child.next = newChild
	}
}

// Busca um valor baseado na chave
func (tree *BPlusTree) search(key string) (string, bool) {
	return searchNode(tree.root, key)
}

// Função auxiliar para buscar um valor em um nó
func searchNode(node *Node, key string) (string, bool) {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}
	if i < len(node.keys) && key == node.keys[i] {
		return node.values[i], true
	}
	if node.isLeaf {
		return "", false
	}
	return searchNode(node.children[i], key)
}

// Imprime a árvore B+
func (tree *BPlusTree) print() {
	printNode(tree.root, 0)
}

// Função auxiliar para imprimir um nó
func printNode(node *Node, level int) {
	fmt.Printf("%sLevel %d %v\n", string(rune(level*4)), level, node.keys)
	if node.isLeaf {
		fmt.Printf("%sValues %v\n", string(rune(level*4)), node.values)
	}
	for _, child := range node.children {
		printNode(child, level+1)
	}
}

// Função principal para testar a árvore B+
func main() {
	t := 3 // Grau mínimo
	bptree := newBPlusTree(t)

	// Inserindo pares chave-valor (nome, sobrenome)
	pairs := map[string]string{
		"John":    "Doe",
		"Jane":    "Smith",
		"Michael": "Johnson",
		"Emily":   "Davis",
		"Daniel":  "Wilson",
		"Sophia":  "Garcia",
		"David":   "Martinez",
		"Emma":    "Anderson",
	}
	for key, value := range pairs {
		bptree.insert(key, value)
	}

	fmt.Println("Árvore B+:")
	bptree.print()

	// Teste de busca
	name := "John"
	if surname, found := bptree.search(name); found {
		fmt.Printf("Sobrenome de %s é %s\n", name, surname)
	} else {
		fmt.Printf("%s não encontrado\n", name)
	}
}
