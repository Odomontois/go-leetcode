package y2024

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	seg, bound := [2]int{}, [2]int{lower - 1, upper}
	res := int64(0)
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		for b := 0; b < 2; b++ {
			for seg[b] < len(nums) && nums[seg[b]]+x <= bound[b] {
				seg[b]++
			}
		}
		res += int64(seg[1] - seg[0])
		if i >= seg[0] && i < seg[1] {
			res--
		}
	}
	return res / 2
}
