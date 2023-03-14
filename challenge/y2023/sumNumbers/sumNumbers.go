package sumnumbers

import "leetcode/data"

type TreeNode = data.TreeNode

func sumNumbers(root *TreeNode) int {
	return traverse(root, 0)
}

func traverse(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return traverse(root.Left, sum) + traverse(root.Right, sum)
}