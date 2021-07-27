package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	best := 100000000
	db := best - target

	for i, x := range nums {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			s := x + nums[j] + nums[k]
			d := s - target
			if d < 0 {
				d = -d
				j++
			} else if d > 0 {
				k--
			} else {
				return target
			}
			if d < db {
				db = d
				best = s
			}
		}
	}
	return best
}
