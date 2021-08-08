package main

import "sort"

type Node struct {
	children []*Node
	rank     int
	val      int
	blocks   int
	size     int
}

type GetNode interface {
	getNode() *Node
}

type Elem struct {
	link GetNode
}

func (node *Node) getNode() *Node {
	return node
}

func (elem *Elem) getNode() *Node {
	root := elem.root()
	if root.link == nil {
		root.link = new(Node)
	}
	return root.link.getNode()
}

func (elem *Elem) root() *Elem {
	if elem.link == nil {
		return elem
	}
	if _, ok := elem.link.(*Node); ok {
		return elem
	}
	root := elem.link.(*Elem).root()
	elem.link = root
	return root
}

func (elem *Elem) merge(el2 *Elem) *Elem {
	r1, r2 := elem.root(), el2.root()
	if r1 == r2 {
		return r1
	}
	s1, s2 := &r1.getNode().size, &r2.getNode().size
	total := *s1 + *s2
	if *s1 > *s2 {
		r2.link, *s1 = r1, total
		return r1
	}
	r1.link, *s2 = r2, total
	return r2

}

const (
	ROW int = iota
	COL
)

type ElemMap = map[int]*Elem

func matrixRankTransform(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])

	elems := make([][]*Elem, n)

	nums := [2][]ElemMap{make([]ElemMap, n), make([]ElemMap, m)}

	for _, ns := range nums {
		for i := range ns {
			ns[i] = make(ElemMap)
		}
	}

	numElem := func(dim, i, val int) *Elem {
		elem := nums[dim][i][val]
		if elem == nil {
			elem = new(Elem)
			elem.getNode().val = val
			nums[dim][i][val] = elem
		}
		return elem
	}

	for i, row := range matrix {
		elems[i] = make([]*Elem, m)
		for j, val := range row {
			elems[i][j] = numElem(ROW, i, val).merge(numElem(COL, j, val))
		}
	}

	for _, slc := range nums {
		for _, m := range slc {
			var nodes []*Node
			for _, elem := range m {
				nodes = append(nodes, elem.getNode())
			}
			sort.Slice(nodes, func(i, j int) bool {
				return nodes[i].val < nodes[j].val
			})
			for i := 1; i < len(nodes); i++ {
				nodes[i-1].children = append(nodes[i-1].children, nodes[i])
				nodes[i].blocks++
			}
		}
	}

	var level []*Node
	rank := 1

	for _, row := range elems {
		for _, e := range row {
			node := e.getNode()
			if node.blocks == 0 {
				level = append(level, node)
				node.blocks = -1
			}
		}
	}

	for len(level) > 0 {
		var next []*Node
		for _, node := range level {
			node.rank = rank
			for _, child := range node.children {
				child.blocks--
				if child.blocks == 0 {
					next = append(next, child)
				}
			}
		}
		rank++
		level = next
	}

	res := make([][]int, n)
	for i, row := range elems {
		res[i] = make([]int, m)
		for j, elem := range row {
			res[i][j] = elem.getNode().rank
		}
	}

	return res
}
