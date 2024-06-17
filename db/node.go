package db

type Node struct {
	IsLeaf   bool
	Keys     []byte
	Values   []byte
	Children []*Node
	Next     *Node
}

func NewNode(t int, isLeaf bool) *Node {
	return &Node{
		IsLeaf:   isLeaf,
		Keys:     make([]byte, 0, 2*t-1),
		Values:   make([]byte, 0, 2*t-1),
		Children: make([]*Node, 0, 2*t),
	}
}

// func (n *Node) Find(key []byte) (int, bool) {
// 	for i, k := range n.Keys {
// 		if k == key[0] {
// 			return i, true
// 		}
// 	}

// 	return -1, false
// }

// func (n *Node) InsertNonFull(t int, key, value []byte) {
// 	i := len(n.Keys) - 1

// 	if n.IsLeaf {
// 		n.Keys = append(n.Keys, 0)
// 		n.Values = append(n.Values, 0)

// 		for i >= 0 && key[0] < n.Keys[i] {
// 			n.Keys[i+1] = n.Keys[i]
// 			n.Values[i+1] = n.Values[i]
// 			i--
// 		}

// 		n.Keys[i+1] = key[0]
// 		n.Values[i+1] = value[0]
// 	} else {
// 		for i >= 0 && key[0] < n.Keys[i] {
// 			i--
// 		}

// 		i++

// 		if len(n.Children[i].Keys) == 2*t-1 {
// 			n.SplitChild(t, i, n.Children[i])

// 			if key[0] > n.Keys[i] {
// 				i++
// 			}
// 		}

// 		n.Children[i].InsertNonFull(t, key, value)
// 	}
// }

// func (n *Node) SplitChild(t, i int, y *Node) {
// 	z := NewNode(t, y.IsLeaf)
// 	z.Keys = append(z.Keys, y.Keys[t:]...)
// 	z.Values = append(z.Values, y.Values[t:]...)

// 	y.Keys = y.Keys[:t-1]
// 	y.Values = y.Values[:t-1]

// 	if !y.IsLeaf {
// 		z.Children = append(z.Children, y.Children[t:]...)
// 		y.Children = y.Children[:t]
// 	}

// 	n.Children = append(n.Children, nil)
// 	copy(n.Children[i+2:], n.Children[i+1:])
// 	n.Children[i+1] = z

// 	n.Keys = append(n.Keys, 0)
// 	copy(n.Keys[i+1:], n.Keys[i:])
// 	n.Keys[i] = y.Keys[t-1]

// 	n.Values = append(n.Values, 0)
// 	copy(n.Values[i+1:], n.Values[i:])
// 	n.Values[i] = y.Values[t-1]
// }

// func (n *Node) Remove(t int, key []byte) {
// 	i, found := n.Find(key)

// 	if found {
// 		if n.IsLeaf {
// 			n.RemoveFromLeaf(i)
// 		} else {
// 			n.RemoveFromNonLeaf(t, i)
// 		}
// 	} else {
// 		if n.IsLeaf {
// 			return
// 		}

// 		var flag int

// 		if i == len(n.Keys) {
// 			flag = i - 1
// 		} else {
// 			flag = i
// 		}

// 		if len(n.Children[flag].Keys) < t {
// 			n.Fill(t, flag)
// 		}

// 		n.Children[flag].Remove(t, key)
// 	}
// }

// func (n *Node) RemoveFromLeaf(i int) {
// 	copy(n.Keys[i:], n.Keys[i+1:])
// 	n.Keys = n.Keys[:len(n.Keys)-1]

// 	copy(n.Values[i:], n.Values[i+1:])
// 	n.Values = n.Values[:len(n.Values)-1]
// }

// func (n *Node) RemoveFromNonLeaf(t, i int) {
// 	k := n.Keys[i]

// 	if len(n.Children[i].Keys) >= t {
// 		pred := n.GetPred(i)
// 		n.Keys[i] = pred
// 		n.Children[i].Remove(t, []byte{pred})
// 	} else if len(n.Children[i+1].Keys) >= t {
// 		succ := n.GetSucc(i)
// 		n.Keys[i] = succ
// 		n.Children[i+1].Remove(t, []byte{succ})
// 	} else {
// 		n.Merge(t, i)
// 		n.Children[i].Remove(t, []byte{k})
// 	}
// }

// func (n *Node) GetPred(i int) byte {
// 	cur := n.Children[i]

// 	for !cur.IsLeaf {
// 		cur = cur.Children[len(cur.Children)-1]
// 	}

// 	return cur.Keys[len(cur.Keys)-1]
// }

// func (n *Node) GetSucc(i int) byte {
// 	cur := n.Children[i+1]

// 	for !cur.IsLeaf {
// 		cur = cur.Children[0]
// 	}

// 	return cur.Keys[0]
// }

// func (n *Node) Fill(t, i int) {
// 	if i != 0 && len(n.Children[i-1].Keys) >= t {
// 		n.BorrowFromPrev(i)
// 	} else if i != len(n.Keys) && len(n.Children[i+1].Keys) >= t {
// 		n.BorrowFromNext(i)
// 	} else {
// 		if i != len(n.Keys) {
// 			n.Merge(t, i)
// 		} else {
// 			n.Merge(t, i-1)
// 		}
// 	}
// }

// func (n *Node) BorrowFromPrev(i int) {
// 	child := n.Children[i]
// 	sibling := n.Children[i-1]

// 	child.Keys = append([]byte{n.Keys[i-1]}, child.Keys...)
// 	child.Values = append([]byte{n.Values[i-1]}, child.Values...)

// 	if !child.IsLeaf {
// 		child.Children = append([]*Node{sibling.Children[len(sibling.Children)-1]}, child.Children...)
// 		sibling.Children = sibling.Children[:len(sibling.Children)-1]
// 	}

// 	n.Keys[i-1] = sibling.Keys[len(sibling.Keys)-1]
// 	n.Values[i-1] = sibling.Values[len(sibling.Values)-1]

// 	sibling.Keys = sibling.Keys[:len(sibling.Keys)-1]
// 	sibling.Values = sibling.Values[:len(sibling.Values)-1]
// }

// func (n *Node) BorrowFromNext(i int) {
// 	child := n.Children[i]
// 	sibling := n.Children[i+1]

// 	child.Keys = append(child.Keys, n.Keys[i])
// 	child.Values = append(child.Values, n.Values[i])

// 	if !child.IsLeaf {
// 		child.Children = append(child.Children, sibling.Children[0])
// 		sibling.Children = sibling.Children[1:]
// 	}

// 	n.Keys[i] = sibling.Keys[0]
// 	n.Values[i] = sibling.Values[0]

// 	sibling.Keys = sibling.Keys[1:]
// 	sibling.Values = sibling.Values[1:]
// }

// func (n *Node) Merge(t, i int) {
// 	child := n.Children[i]
// 	sibling := n.Children[i+1]

// 	child.Keys = append(child.Keys, n.Keys[i])
// 	child.Values = append(child.Values, n.Values[i])

// 	child.Keys = append(child.Keys, sibling.Keys...)
// 	child.Values = append(child.Values, sibling.Values...)

// 	if !child.IsLeaf {
// 		child.Children = append(child.Children, sibling.Children...)
// 	}

// 	n.Keys = append(n.Keys[:i], n.Keys[i+1:]...)
// 	n.Values = append(n.Values[:i], n.Values[i+1:]...)

// 	n.Children = append(n.Children[:i+1], n.Children[i+2:]...)
// }

// func (n *Node) Traverse() []byte {
// 	var result []byte

// 	for i, child := range n.Children {
// 		if i > 0 {
// 			result = append(result, child.Traverse()...)
// 		}

// 		if i < len(n.Keys) {
// 			result = append(result, n.Keys[i])
// 		}
// 	}

// 	if len(n.Children) > 0 {
// 		result = append(result, n.Children[len(n.Children)-1].Traverse()...)
// 	}

// 	return result
// }

// func (n *Node) Search(key []byte) (string, bool) {
// 	i, found := n.Find(key)

// 	if found {
// 		return string(n.Values[i]), true
// 	}

// 	if n.IsLeaf {
// 		return "", false
// 	}

// 	return n.Children[i].Search(key)
// }

// func (n *Node) Update(key, value []byte) {
// 	i, found := n.Find(key)

// 	if found {
// 		n.Values[i] = value[0]
// 		return
// 	}

// 	if n.IsLeaf {
// 		return
// 	}

// 	n.Children[i].Update(key, value)
// }

// func (n *Node) Get(t int, key []byte) (string, bool) {
// 	i, found := n.Find(key)

// 	if found {
// 		return string(n.Values[i]), true
// 	}

// 	if n.IsLeaf {
// 		return "", false
// 	}

// 	return n.Children[i].Get(t, key)
// }

// func (n *Node) GetRange(t int, key []byte, result []string) []string {
// 	i := 0

// 	for i < len(n.Keys) && key[0] > n.Keys[i] {
// 		i++
// 	}

// 	if i < len(n.Keys) && key[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i].GetRange(t, key, result)
// 	}

// 	return result
// }

// func (n *Node) GetRangeAll(t int, result []string) []string {
// 	i := 0

// 	for i < len(n.Keys) {
// 		result = append(result, string(n.Values[i]))
// 		i++
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i].GetRangeAll(t, result)
// 	}

// 	return result
// }

// func (n *Node) GetRangeBetween(t int, key1, key2 []byte, result []string) []string {
// 	i := 0

// 	for i < len(n.Keys) && key1[0] > n.Keys[i] {
// 		i++
// 	}

// 	if i < len(n.Keys) && key1[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i].GetRangeBetween(t, key1, key2, result)
// 	}

// 	if i < len(n.Keys) && key2[0] > n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	return result
// }

// func (n *Node) GetRangeBetweenAll(t int, key1, key2 []byte, result []string) []string {
// 	i := 0

// 	for i < len(n.Keys) {
// 		result = append(result, string(n.Values[i]))
// 		i++
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i].GetRangeBetweenAll(t, key1, key2, result)
// 	}

// 	return result
// }

// func (n *Node) GetRangeReverse(t int, key []byte, result []string) []string {
// 	i := len(n.Keys) - 1

// 	for i >= 0 && key[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetRangeReverse(t, key, result)
// 	}

// 	return result
// }

// func (n *Node) GetRangeAllReverse(t int, result []string) []string {
// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		result = append(result, string(n.Values[i]))
// 		i--
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetRangeAllReverse(t, result)
// 	}

// 	return result
// }

// func (n *Node) GetRangeBetweenReverse(t int, key1, key2 []byte, result []string) []string {
// 	i := len(n.Keys) - 1

// 	for i >= 0 && key1[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key1[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetRangeBetweenReverse(t, key1, key2, result)
// 	}

// 	if i >= 0 && key2[0] < n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	return result
// }

// func (n *Node) GetRangeBetweenAllReverse(t int, key1, key2 []byte, result []string) []string {
// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		result = append(result, string(n.Values[i]))
// 		i--
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetRangeBetweenAllReverse(t, key1, key2, result)
// 	}

// 	return result
// }

// func (n *Node) GetReverse(t int, key []byte) (string, bool) {
// 	i, found := n.Find(key)

// 	if found {
// 		return string(n.Values[i]), true
// 	}

// 	if n.IsLeaf {
// 		return "", false
// 	}

// 	return n.Children[i+1].GetReverse(t, key)
// }

// func (n *Node) GetReverseAll(t int) []string {
// 	var result []string

// 	if !n.IsLeaf {
// 		result = n.Children[len(n.Children)-1].GetReverseAll(t)
// 	}

// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		result = append(result, string(n.Values[i]))
// 		i--
// 	}

// 	return result
// }

// func (n *Node) GetReverseBetween(t int, key1, key2 []byte) []string {
// 	var result []string

// 	i := len(n.Keys) - 1

// 	for i >= 0 && key1[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key1[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetReverseBetween(t, key1, key2)
// 	}

// 	if i >= 0 && key2[0] < n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	return result
// }

// func (n *Node) GetReverseBetweenAll(t int, key1, key2 []byte) []string {
// 	var result []string

// 	if !n.IsLeaf {
// 		result = n.Children[len(n.Children)-1].GetReverseBetweenAll(t, key1, key2)
// 	}

// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		result = append(result, string(n.Values[i]))
// 		i--
// 	}

// 	return result
// }

// func (n *Node) UpdateReverse(t int, key, value []byte) {
// 	i, found := n.Find(key)

// 	if found {
// 		n.Values[i] = value[0]
// 		return
// 	}

// 	if n.IsLeaf {
// 		return
// 	}

// 	n.Children[i+1].UpdateReverse(t, key, value)
// }

// func (n *Node) GetReverseRange(t int, key []byte, result []string) []string {
// 	i := len(n.Keys) - 1

// 	for i >= 0 && key[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetReverseRange(t, key, result)
// 	}

// 	return result
// }

// func (n *Node) GetReverseRangeAll(t int, result []string) []string {
// 	if !n.IsLeaf {
// 		result = n.Children[len(n.Children)-1].GetReverseRangeAll(t, result)
// 	}

// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		result = append(result, string(n.Values[i]))
// 		i--
// 	}

// 	return result
// }

// func (n *Node) GetReverseRangeBetween(t int, key1, key2 []byte, result []string) []string {
// 	i := len(n.Keys) - 1

// 	for i >= 0 && key1[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key1[0] == n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	if !n.IsLeaf {
// 		result = n.Children[i+1].GetReverseRangeBetween(t, key1, key2, result)
// 	}

// 	if i >= 0 && key2[0] < n.Keys[i] {
// 		result = append(result, string(n.Values[i]))
// 	}

// 	return result
// }

// func (n *Node) GetReverseRangeBetweenAll(t int, key1, key2 []byte, result []string) []string {
// 	if !n.IsLeaf {
// 		result = n.Children[len(n.Children)-1].GetReverseRangeBetweenAll(t, key1, key2, result)
// 	}

// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		result = append(result, string(n.Values[i]))
// 		i--
// 	}

// 	return result
// }

// func (n *Node) UpdateReverseRange(t int, key, value []byte) {
// 	i := len(n.Keys) - 1

// 	for i >= 0 && key[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key[0] == n.Keys[i] {
// 		n.Values[i] = value[0]
// 		return
// 	}

// 	if !n.IsLeaf {
// 		n.Children[i+1].UpdateReverseRange(t, key, value)
// 	}
// }

// func (n *Node) UpdateReverseRangeAll(t int, value []byte) {
// 	if !n.IsLeaf {
// 		n.Children[len(n.Children)-1].UpdateReverseRangeAll(t, value)
// 	}

// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		n.Values[i] = value[0]
// 		i--
// 	}
// }

// func (n *Node) UpdateReverseRangeBetween(t int, key1, key2, value []byte) {
// 	i := len(n.Keys) - 1

// 	for i >= 0 && key1[0] < n.Keys[i] {
// 		i--
// 	}

// 	if i >= 0 && key1[0] == n.Keys[i] {
// 		n.Values[i] = value[0]
// 	}

// 	if !n.IsLeaf {
// 		n.Children[i+1].UpdateReverseRangeBetween(t, key1, key2, value)
// 	}

// 	if i >= 0 && key2[0] < n.Keys[i] {
// 		n.Values[i] = value[0]
// 	}
// }

// func (n *Node) UpdateReverseRangeBetweenAll(t int, key1, key2, value []byte) {
// 	if !n.IsLeaf {
// 		n.Children[len(n.Children)-1].UpdateReverseRangeBetweenAll(t, key1, key2, value)
// 	}

// 	i := len(n.Keys) - 1

// 	for i >= 0 {
// 		n.Values[i] = value[0]
// 		i--
// 	}
// }
