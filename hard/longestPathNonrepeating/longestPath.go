// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/

package longestpath

func longestPath(parent []int, s string) int {
	tree := makeTree(parent, s)
	path, _ := tree.longest(0)
	return path
}

type bytee interface {
	~string | ~[]byte
}

type tree[Vals bytee] struct {
	vals     Vals
	children [][]int
}

func makeTree[Vals bytee](parents []int, vals Vals) tree[Vals] {
	children := make([][]int, len(parents))
	for i, p := range parents {
		if i == 0 {
			continue
		}
		parent := &children[p]
		*parent = append(*parent, i)
	}
	return tree[Vals]{vals, children}
}

func setMax(best, second *int, val int) {
	if *best < val {
		*second, *best = *best, val
	} else if *second < val {
		*second = val
	}
}

func (tree *tree[Vals]) longest(i int) (path int, way int) {
	var dummy, second int
	for _, child := range tree.children[i] {
		cpath, cway := tree.longest(child)
		if tree.vals[i] != tree.vals[child] {
			setMax(&way, &second, cway)
		}
		setMax(&path, &dummy, cpath)
	}
	way++
	setMax(&path, &dummy, way+second)
	return
}
