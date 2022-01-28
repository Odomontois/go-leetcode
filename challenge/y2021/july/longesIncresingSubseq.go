package main

import "sort"

func lengthOfLIS(nums []int) int {
	ends := []int{}
	for _, x := range nums {
		i := sort.SearchInts(ends, x)
		if i == len(ends) {
			ends = append(ends, x)
		} else {
			ends[i] = x
		}
	}
	return len(ends)
}
