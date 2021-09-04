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
	res := make([]int, n)
	tree.fillDists(res)
	return res
}

func (t *Tree) fillDists(dists []int) {
	dists[t.idx] = t.sum
	for _, child := range t.children {
		child.descentFillDists(dists, t.sum)
	}
}

func (t *Tree) descentFillDists(dists []int, parentSum int) {
	dists[t.idx] = parentSum - 2*t.count + len(dists) - 2
	for _, child := range t.children {
		child.descentFillDists(dists, dists[t.idx])
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
		child := buildTree(child, start, edges)
		tree.children = append(tree.children, child)
		tree.sum += child.sum + child.count + 1
		tree.count += child.count + 1
	}
	return &tree
}
