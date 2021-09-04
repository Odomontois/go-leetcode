package sumtreedist

func SumOfDistancesInTree(n int, edges [][]int) []int {
	return sumOfDistancesInTree(n, edges)
}

type Tree struct {
	idx, sum, count int
	children        []*Tree
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	tree := buildTree(0, -1, buildEdgeMap(n, edges))
	tree.fillSubTreeDists()
	res := make([]int, n)
	tree.descentFillDists(res, 0, 0)
	return res
}

func (t *Tree) descentFillDists(dists []int, parentSum, parentCount int) {
	dists[t.idx] = parentSum + t.sum
	for _, child := range t.children {
		curCount := parentCount + t.count - child.count
		curSum := parentSum + t.sum - child.sum - child.count + curCount - 1
		child.descentFillDists(dists, curSum, curCount)
	}
}

func (t *Tree) fillSubTreeDists() {
	for _, child := range t.children {
		child.fillSubTreeDists()
		t.sum += child.sum + child.count + 1
		t.count += child.count + 1
	}
}

func buildEdgeMap(n int, edges [][]int) [][]int {
	edgeMap := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		edgeMap[x] = append(edgeMap[x], y)
		edgeMap[y] = append(edgeMap[y], x)
	}
	return edgeMap
}

func buildTree(start int, parent int, edges [][]int) *Tree {
	tree := Tree{idx: start}
	for _, child := range edges[start] {
		if child == parent {
			continue
		}
		tree.children = append(tree.children, buildTree(child, start, edges))
	}
	return &tree
}
