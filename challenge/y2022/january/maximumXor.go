package main

func findMaximumXOR(nums []int) int {
	mask := 0
	for check := 1 << 31; check != 0; check >>= 1 {
		if existsDiff(nums, mask|check) {
			mask |= check
		}
	}
	return mask
}

func existsDiff(nums []int, mask int) bool {
	bits := make(map[int]bool)
	for _, num := range nums {
		masked := num & mask
		invert := masked ^ mask
		if has := bits[invert]; has {
			return true
		}
		bits[masked] = true
	}
	return false
}
