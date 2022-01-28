package main

func findPeakElement(nums []int) int {
	if len(nums) < 2 || nums[0] > nums[1] {
		return 0
	}
	n := len(nums) - 1
	if nums[n] > nums[n-1] {
		return n
	}
	m := n / 2
	if nums[m] > nums[m+1] {
		return findPeakElement(nums[:m+1])
	}
	return findPeakElement(nums[m:]) + m
}
