package removalbintreequeries

import "leetcode/data"

type TreeNode = data.TreeNode

func treeQueries(root *TreeNode, queries []int) (res []int) {
	trav := makeTraverse(root)
	for _, q := range queries {
		res = append(res, trav.calc(q))
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Traverse struct {
	interval map[int][2]int
	values   []int
	maxTree  []int
}

func makeTraverse(root *TreeNode) (result Traverse) {
	result.interval = make(map[int][2]int)
	result.recur(root, 0)
	result.maxTree = make([]int, len(result.values)*4)
	result.fillMaxTree(0, len(result.values), 0)
	return
}

func (trav *Traverse) recur(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	start := len(trav.values)
	trav.recur(node.Left, depth+1)
	trav.values = append(trav.values, depth)
	trav.recur(node.Right, depth+1)
	end := len(trav.values)
	trav.interval[node.Val] = [2]int{start, end}
}

func (trav *Traverse) fillMaxTree(start, end, idx int) {
	if end-start == 1 {
		trav.maxTree[idx] = trav.values[start]
		return
	}
	m, l, r := (start+end+1)/2, idx*2+1, idx*2+2

	trav.fillMaxTree(start, m, l)
	trav.fillMaxTree(m, end, r)
	trav.maxTree[idx] = max(trav.maxTree[l], trav.maxTree[r])
}

func (trav *Traverse) calc(val int) int {
	in, ok := trav.interval[val]
	if !ok {
		return trav.maxTree[0]
	}
	n := len(trav.values)
	before := trav.calcIter(0, in[0], 0, n, 0)
	after := trav.calcIter(in[1], len(trav.values), 0, n, 0)
	return max(before, after)
}

func (trav *Traverse) calcIter(from, to, start, end, idx int) int {
	if from <= start && end <= to {
		return trav.maxTree[idx]
	}
	if to <= start || end <= from || end-start == 1 {
		return -1
	}

	m, l, r := (start+end+1)/2, idx*2+1, idx*2+2
	lm := trav.calcIter(from, to, start, m, l)
	rm := trav.calcIter(from, to, m, end, r)
	return max(lm, rm)
}
