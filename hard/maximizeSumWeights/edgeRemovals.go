package maximizesumweights

import (
	"slices"
)

func maximizeSumOfWeights(edges [][]int, k int) int64 {
	t := newTree(edges)
	return t.dfs(0, -1, k).exclusive
}

type edge struct{ vertix, weight int }
type tree [][]edge

func newTree(edges [][]int) tree {
	t := make(tree, len(edges)+1)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		t[x] = append(t[x], edge{y, w})
		t[y] = append(t[y], edge{x, w})
	}
	return t
}

func (t tree) dfs(i, p, k int) (result struct{ exclusive, inclusive int64 }) {
	var diffs []int64
	for _, e := range t[i] {
		if e.vertix == p {
			continue
		}
		res := t.dfs(e.vertix, i, k)
		result.inclusive += res.exclusive
		diffs = append(diffs, res.inclusive-res.exclusive+int64(e.weight))
	}
	result.exclusive = result.inclusive
	slices.Sort(diffs)
	slices.Reverse(diffs)
	for i, d := range diffs {
		if d < 0 || i == k {
			break
		}
		result.exclusive += d
		if i < k-1 {
			result.inclusive += d
		}
	}
	return
}
