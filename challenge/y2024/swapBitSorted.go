package y2024

import (
	"math/bits"
	"slices"
)

func oneEq(x, y int) bool {
	return bits.OnesCount(uint(x)) == bits.OnesCount(uint(y))
}
func canSortArray(nums []int) bool {
	nums2 := slices.Clone(nums)
	slices.Sort(nums2)
	start := 0
	for i := range nums {
		if !oneEq(nums[i], nums[start]) {
			slices.Sort(nums[start:i])
			start = i
		}
	}
	slices.Sort(nums[start:])
	return slices.Equal(nums, nums2)
}
