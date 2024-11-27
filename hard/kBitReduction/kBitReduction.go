package kbitreduction

import "math/bits"

//3352
// https://leetcode.com/problems/count-k-reducible-numbers-less-than-n/

func countKReducibleNumbers(s string, k int) int {
	panic("not implemented")
}

func stepsReduction(k int) (res int) {
	for k > 1 {
		k = bits.OnesCount(uint(k))
		res++
	}
	return res
}


