package main

import (
	"leetcode/data"
	"math"
	"sync"
)

type TreeNode = data.TreeNode

func goodNodes(root *TreeNode) int {
	var wg sync.WaitGroup
	var rec func(*TreeNode, int)
	var res int

	look := func(node *TreeNode, mx int) {
		if node == nil {
			return
		}
		wg.Add(1)
		go rec(node, mx)
	}
	rec = func(node *TreeNode, mx int) {
		defer wg.Done()
		if node.Val >= mx {
			res++
			mx = node.Val
		}
		wg.Add(2)
		look(node.Left, mx)
		look(node.Right, mx)
	}
	look(root, math.MinInt32)
	wg.Wait()
	return res
}
