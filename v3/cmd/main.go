package main

import (
	"encoding/binary"
	"fmt"
)

const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

func assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}

func main() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE

	assert(node1max <= BTREE_PAGE_SIZE)

	n := make(Node, BTREE_PAGE_SIZE)

	n.setHeader(2, 5)

	fmt.Println(n.btype(), n.nkeys())

	n.setPtr(0, 1)

	fmt.Println(n.getPtr(1))
}

type Node []byte

func (node Node) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}

func (node Node) btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node Node) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node Node) getPtr(idx uint16) uint64 {
	assert(idx < node.nkeys())
	pos := HEADER + 8*idx
	return binary.LittleEndian.Uint64(node[pos:])
}
func (node Node) setPtr(idx uint16, val uint64)

//  {
// 	assert(idx < node.nkeys())
// 	pos := HEADER + 8*idx
// 	binary.LittleEndian.PutUint64(node[pos:], val)
// }

type Tree struct {
	root uint64
	get  func(uint64) []byte
	new  func([]byte) uint64
	del  func(uint64)
}
