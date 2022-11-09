package main

import "math"

type tree struct {
	val      int
	children []*tree
}

func build(nums []int, edgeLists [][]int, i, p int) *tree {
	root := tree{val: nums[i]}
	for _, c := range edgeLists[i] {
		if c != p {
			root.children = append(root.children, build(nums, edgeLists, c, i))
		}
	}
	return &root
}

func (root *tree) trySplit(n int) (sum int, found bool) {
	sum += root.val
	for _, t := range root.children {
		if s, ok := t.trySplit(n); ok {
			sum += s
		} else {
			return
		}
	}
	if sum == n {
		sum = 0
	}
	return sum, sum < n
}

func componentValue(nums []int, edges [][]int) int {
	edgeLists := make([][]int, len(nums))
	for _, e := range edges {
		edgeLists[e[0]] = append(edgeLists[e[0]], e[1])
		edgeLists[e[1]] = append(edgeLists[e[1]], e[0])
	}
	root := build(nums, edgeLists, 0, -1)
	total, _ := root.trySplit(math.MaxInt)
	var divs []int
	for i := 1; i*i <= total; i++ {
		if total%i == 0 {
			divs = append(divs, i)
		}
	}
	for i := len(divs) - 1; i >= 0; i-- {
		x := divs[i]
		if x*x != total {
			divs = append(divs, total/x)
		}
	}
	for _, x := range divs {
		if s, ok := root.trySplit(x); ok && s == 0 {
			return total/x - 1
		}
	}

	return 0
}
