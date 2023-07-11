package distancek

import "leetcode/data"

type TreeNode = data.TreeNode

func distanceK(root *TreeNode, target *TreeNode, k int) (res []int) {
	var path []Path
	findPath(root, target, &path)
	descendPath(root, path, &res, k-len(path))
	return
}

type Path uint8

const (
	L Path = iota
	R
)

func findPath(root *TreeNode, target *TreeNode, path *[]Path) bool {
	if root == target {
		return true
	}
	if root == nil {
		return false
	}
	n := len(*path)
	*path = append(*path, L)
	if findPath(root.Left, target, path) {
		return true
	}
	(*path)[n] = R
	if findPath(root.Right, target, path) {
		return true
	}
	*path = (*path)[:n]
	return false
}

func descendPath(root *TreeNode, path []Path, res *[]int, k int) {
	if len(path) == 0 {
		collectDown(root, res, k)
		return
	}
	if k == 0 {
		*res = append(*res, root.Val)
	}
	var straight, side *TreeNode
	switch path[0] {
	case L:
		straight, side = root.Left, root.Right
	case R:
		straight, side = root.Right, root.Left
	}

	descendPath(straight, path[1:], res, k+1)
	if k > 0 {
		collectDown(side, res, k-1)
	}
}

func collectDown(root *data.TreeNode, res *[]int, k int) {
	if root == nil {
		return
	}
	if k == 0 {
		*res = append(*res, root.Val)
		return
	}
	collectDown(root.Left, res, k-1)
	collectDown(root.Right, res, k-1)
}
