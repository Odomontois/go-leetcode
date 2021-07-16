package main

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i, x := range nums {
		for j, y := range nums[i+1:] {
			res += sort.SearchInts(nums[i+j+2:], x+y)
		}
	}
	return res
}
