package y2024

import (
	"leetcode/data"
)

type ListNode = data.ListNode

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := range m {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	lim, xy, dxy := [2]int{m - 1, n}, [2]int{0, -1}, [2]int{1, 1}
outer:
	for {
		for _, dim := range [2]int{1, 0} {
			if lim[dim] == 0 {
				break outer
			}
			for i := 0; i < lim[dim]; i++ {
				if head == nil {
					break outer
				}
				xy[dim] += dxy[dim]
				res[xy[0]][xy[1]] = head.Val
				head = head.Next
			}
			dxy[dim] = -dxy[dim]
			lim[dim]--
		}
	}
	return res
}
