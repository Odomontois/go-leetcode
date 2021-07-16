package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	hm := make(map[int]bool)
	idx := make(map[[4]int]bool)
	var res [][]int
	sort.Ints(nums)
	for _, x := range nums {
		hm[x] = true
	}
	for i, x := range nums {
		for j, y := range nums[i+1:] {
			for k, z := range nums[i+j+2:] {
				d := i + j + k + 3
				w := target - x - y - z
				if d < len(nums) && w >= nums[d] && hm[w] {
					idx[[4]int{x, y, z, w}] = true
				}
			}
		}
	}
	for k := range idx {
		it := make([]int, 4)
		copy(it, k[:])
		res = append(res, it)
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
