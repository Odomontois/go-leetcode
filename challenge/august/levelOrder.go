package main

import "leetcode/data"

type Node = data.Node

func levelOrder(root *Node) (res [][]int) {
	cur := []*Node{root}
	for root != nil && len(cur) > 0 {
		next, level := []*Node{}, []int{}
		for _, n := range cur {
			level, next = append(level, n.Val), append(next, n.Children...)
		}
		res, cur = append(res, level), next
	}
	return res
}
