package main

import "leetcode/data"

type TreeNode = data.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums) / 2
	return &TreeNode{Val: nums[n], Left: sortedArrayToBST(nums[:n]), Right: sortedArrayToBST(nums[n+1:])}
}
