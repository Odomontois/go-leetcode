package main

func findLength(nums1 []int, nums2 []int) int {
	prev := make([]int, len(nums2)+1)
	best := 0
	for _, x := range nums1 {
		cur := make([]int, len(nums2)+1)
		for j, y := range nums2 {
			if x == y {
				cur[j+1] = prev[j] + 1
				if cur[j+1] > best {
					best = cur[j+1]
				}
			}
		}
		prev = cur
	}
	return best
}
