package minnumbertaps

const Inf int = 1_000_000_000

func minTaps(n int, ranges []int) int {
	segTree := MakeSegTree(n + 1)
	segTree.Update(0, 0)
	for i, r := range ranges {
		from, to := i-r, i+r
		if to > n {
			to = n
		}
		if from <= 0 {
			segTree.Update(to, 1)
		} else {
			best := 1 + segTree.RangeMin(from, to)
			segTree.Update(to, best)
		}

	}
	if x := segTree.RangeMin(n, n+1); x >= Inf {
		return -1
	} else {
		return x
	}
}

type SegTree []int

func MakeSegTree(size int) SegTree {
	res := make(SegTree, size*4)
	for i := range res {
		res[i] = Inf
	}
	return res
}

func (segTree SegTree) Update(i int, x int) {
	start, end, node := 0, len(segTree)/4, 0
	for {
		if segTree[node] > x {
			segTree[node] = x
		}
		if end-start == 1 {
			break
		}
		m := start + (end-start)/2
		if i < m {
			node, end = node*2+1, m
		} else {
			node, start = node*2+2, m
		}
	}
}

func (segTree SegTree) RangeMin(from, to int) int {
	var iter func(int, int, int) int
	iter = func(node, start, end int) int {
		if start >= from && end <= to {
			return segTree[node]
		}
		if start >= to || end <= from || end-start == 1 {
			return Inf
		}
		m := start + (end-start)/2
		left := iter(node*2+1, start, m)
		if right := iter(node*2+2, m, end); right < left {
			return right
		}
		return left
	}

	return iter(0, 0, len(segTree)/4)
}
