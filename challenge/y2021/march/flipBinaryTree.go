package march

import (
	"leetcode/data"
)

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var walk func(*TreeNode) bool
	flips := []int{}
	k := 0
	walk = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if voyage[k] != node.Val {
			return false
		}
		if node.Left != nil && k+1 < len(voyage) && node.Left.Val != voyage[k+1] {
			flips = append(flips, node.Val)
			node.Left, node.Right = node.Right, node.Left
		}
		k++
		return walk(node.Left) && walk(node.Right)
	}
	if !walk(root) {
		return []int{-1}
	}
	return flips
}

type TreeNode = data.TreeNode
