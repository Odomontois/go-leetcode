package y2024

func countPairs(root *TreeNode, distance int) int {
	return leaves(root, distance).good
}

type leavesRes struct {
	dists []int
	good  int
}

func (l leavesRes) merge(r leavesRes, dist int) leavesRes {
	ln, rn := len(l.dists), len(r.dists)
	if ln < rn {
		l, r, ln, rn = r, l, rn, ln
	}
	l.good += r.good
	u := rn
	if dist-1 < u {
		u = dist - 1
	}

	for i := 0; i < u; i++ {
		rd := r.dists[rn-i-1]
		if i+2 <= rn {
			rd -= r.dists[rn-i-2]
		}
		j := ln + i - dist
		ld := l.dists[ln-1]
		if j >= 0 {
			ld -= l.dists[j]
		}
		l.good += ld * rd
	}
	for i := 0; i < rn; i++ {
		l.dists[ln-i-1] += r.dists[rn-i-1]
	}
	l.dists = append(l.dists, l.dists[ln-1])
	return l
}

func leaves(node *TreeNode, goodDistance int) leavesRes {
	if node == nil {
		return leavesRes{}
	}
	if node.Left == nil && node.Right == nil {
		return leavesRes{[]int{1}, 0}
	}
	left := leaves(node.Left, goodDistance)
	right := leaves(node.Right, goodDistance)
	return left.merge(right, goodDistance)
}
