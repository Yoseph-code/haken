package main

import (
	"flag"
	"log"

	"github.com/Yoseph-code/haken/server"
)

func main() {
	flag.Uint("p", server.DefaultListenAddr, "HTTP network address")

	flag.Parse()

	addr := flag.Lookup("p").Value.(flag.Getter).Get().(uint)

	s := server.New(server.Config{
		ListenAddr: addr,
	})

	if err := s.Run(); err != nil {
		log.Panic(err)
	}
}

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// type node struct {
// 	value int
// 	left  *node
// 	right *node
// }

// func (n node) String() string {
// 	return strconv.Itoa(n.value)
// }

// type bst struct {
// 	root *node
// 	len  int
// }

// func (b bst) String() string {
// 	sb := strings.Builder{}
// 	b.inOrderTransversal(&sb)
// 	return sb.String()
// }

// func (b bst) inOrderTransversal(sb *strings.Builder) {
// 	b.inOrderTransversalByNode(sb, b.root)
// }

// func (b bst) inOrderTransversalByNode(sb *strings.Builder, root *node) {
// 	if root == nil {
// 		return
// 	}
// 	b.inOrderTransversalByNode(sb, root.left)
// 	sb.WriteString(fmt.Sprintf("%s ", root))
// 	b.inOrderTransversalByNode(sb, root.right)
// }

// func (b bst) add(value int) {
// 	b.root = b.addByNode(b.root, value)
// 	b.len++
// }

// func (b *bst) addByNode(root *node, value int) *node {
// 	if root == nil {
// 		return &node{value: value, left: nil, right: nil}
// 	}

// 	if value < root.value {
// 		root.left = b.addByNode(root.left, value)
// 	} else if value > root.value {
// 		root.right = b.addByNode(root.right, value)
// 	} else {
// 		root.value = value
// 	}

// 	return root
// }

// func (b bst) search(value int) bool {
// 	return b.searchByNode(b.root, value)
// }

// func (b bst) searchByNode(root *node, value int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if value < root.value {
// 		return b.searchByNode(root.left, value)
// 	} else if value > root.value {
// 		return b.searchByNode(root.right, value)
// 	}

// 	return true
// }

// func (b bst) remove(value int) {
// 	b.root = b.removeByNode(b.root, value)
// }

// func (b *bst) removeByNode(root *node, value int) *node {
// 	if root == nil {
// 		return nil
// 	}

// 	if value < root.value {
// 		root.left = b.removeByNode(root.left, value)
// 	} else if value > root.value {
// 		root.right = b.removeByNode(root.right, value)
// 	} else {
// 		if root.left == nil {
// 			return root.right
// 		} else if root.right == nil {
// 			return root.left
// 		}

// 		root.value = b.minValue(root.right)
// 		root.right = b.removeByNode(root.right, root.value)
// 	}

// 	return root
// }

// func (b bst) minValue(root *node) int {
// 	minValue := root.value
// 	for root.left != nil {
// 		minValue = root.left.value
// 		root = root.left
// 	}

// 	return minValue
// }

// func (b bst) size() int {
// 	return b.len
// }

// func (b bst) update(value int) {
// 	b.root = b.updateByNode(b.root, value)
// }

// func (b *bst) updateByNode(root *node, value int) *node {
// 	if root == nil {
// 		return nil
// 	}

// 	if value < root.value {
// 		root.left = b.updateByNode(root.left, value)
// 	} else if value > root.value {
// 		root.right = b.updateByNode(root.right, value)
// 	} else {
// 		root.value = value
// 	}

// 	return root
// }

// func main() {
// 	n := &node{value: 1, left: nil, right: nil}
// 	n.left = &node{value: 0, left: nil, right: nil}
// 	n.right = &node{value: 2, left: nil, right: nil}

// 	b := bst{root: n, len: 3}

// 	fmt.Println(b)

// 	b.add(3)
// 	b.add(10)
// 	b.add(9)
// 	b.add(7)
// 	b.add(4)
// 	b.add(6)
// 	b.add(8)
// 	b.add(5)

// 	fmt.Println(b)

// 	fmt.Println(b.search(50))

// 	b.remove(0)

// 	fmt.Println(b)
// }
