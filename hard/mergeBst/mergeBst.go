package main

import (
	"leetcode/data"
	"math"
)

type TreeNode = data.TreeNode

func checkBst(root *TreeNode, from, to int) bool {
	return root == nil || root.Val > from && root.Val < to &&
		checkBst(root.Left, from, root.Val) &&
		checkBst(root.Right, root.Val, to)
}

func size(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + size(root.Left) + size(root.Right)
}

func canMerge(trees []*TreeNode) *TreeNode {
	roots := make(map[int]*TreeNode)
	seen := make(map[int]bool)
	count := 1

	addLeaf := func(leaf **TreeNode) bool {
		if *leaf == nil {
			return true
		}
		n := (*leaf).Val
		if _, ok := seen[n]; ok {
			return false
		}
		seen[n] = true
		if root := roots[n]; root != nil {
			*leaf = root
			delete(roots, n)
		}
		count++
		return true
	}

	for _, tree := range trees {
		if _, ok := roots[tree.Val]; ok {
			return nil
		}
		roots[tree.Val] = tree
	}

	for _, tree := range trees {
		if !addLeaf(&tree.Left) || !addLeaf(&tree.Right) {
			return nil
		}
	}

	if len(roots) != 1 {
		return nil
	}
	for _, root := range roots {
		if checkBst(root, math.MinInt64, math.MaxInt64) && count == size(root) {
			return root
		}
	}
	return nil
}
