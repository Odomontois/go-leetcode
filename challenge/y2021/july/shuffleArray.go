package main

import "math/rand"

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

/** Resets the array to its original configuration and return it. */
func (sol *Solution) Reset() []int {
	return *sol
}

/** Returns a random shuffling of the array. */
func (sol *Solution) Shuffle() []int {
	res := make([]int, len(*sol))
	copy(res, []int(*sol))
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return res
}
