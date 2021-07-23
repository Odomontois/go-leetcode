package main

import "leetcode/data"

type TreeNode = data.TreeNode

var ZERO TreeNode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if *root == ZERO {
		return nil
	}
	return root
}
