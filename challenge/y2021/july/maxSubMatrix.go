package main

import (
	"github.com/emirpasic/gods/maps/treemap"
)

const MinInt int = (-1 << 63)

func maxSumSubmatrix(matrix [][]int, k int) int {
	colSums := make([][]int, len(matrix))
	for i := range colSums {
		colSums[i] = sliceSums(matrix[i])
	}
	m := len(matrix[0])
	best := MinInt
	ribbon := make([]int, len(matrix))
	for i := 0; i < m; i++ {
		for j := i + 1; j <= m; j++ {
			for u := range ribbon {
				ribbon[u] = colSums[u][j] - colSums[u][i]
			}
			if cur := bestSum(ribbon, k); cur > best {
				best = cur
			}
		}
	}
	return best
}

func sliceSums(is []int) []int {
	sums := make([]int, len(is)+1)
	for j, a := range is {
		sums[j+1] = sums[j] + a
	}
	return sums
}

func bestSum(is []int, k int) int {
	ts := treemap.NewWithIntComparator()
	ts.Put(0, nil)
	best, sum := MinInt, 0
	for _, x := range is {
		sum += x
		if yi, _ := ts.Ceiling(sum - k); yi != nil {
			if cur := sum - yi.(int); cur > best {
				best = cur
			}
		}
		ts.Put(sum, nil)
	}
	return best
}
