package main

import "sort"

func minSetSize(arr []int) int {
	counts := map[int]int{}
	total := 0
	ams := []int{}
	for _, x := range arr {
		counts[x]++
		total++
	}
	for _, v := range counts {
		ams = append(ams, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ams)))
	total /= 2
	for total > 0 {
		total -= ams[0]
		ams = ams[1:]
	}
	return len(counts) - len(ams)
}
