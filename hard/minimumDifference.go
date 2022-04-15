package hard

import (
	"math"
	"sort"
)

func minimumDifference(nums []int) int {
	n := len(nums) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	hsum := (sum + 1) / 2

	best := math.MaxInt
	try := func(res int) {
		if res < best {
			best = res
		}
	}
	lefts, rights := sums(nums[:n]), sums(nums[n:])
	for k, left := range lefts {
		right := rights[n-k]
		for _, lsum := range left {
			rp := sort.SearchInts(right, hsum-lsum)
			if rp > 0 {
				try(sum - 2*(lsum+right[rp-1]))
			}
			if rp < len(right) {
				try(2*(lsum+right[rp]) - sum)
			}
		}
	}
	return best
}

func sums(nums []int) [][]int {
	res := [][]int{{0}}
	for i, num := range nums {
		res = append(res, []int{})
		for j := i; j >= 0; j-- {
			for _, sum := range res[j] {
				res[j+1] = append(res[j+1], sum+num)
			}
		}
	}
	for _, sums := range res {
		sort.Ints(sums)
	}
	return res
}
