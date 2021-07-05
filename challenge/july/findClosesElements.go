package main

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	i := sort.SearchInts(arr, x) - 1
	j := i + 1
	for u := 0; u < k; u++ {
		if i < 0 || j < len(arr) && arr[j]-x < x-arr[i] {
			j++
		} else {
			i--
		}
	}
	return arr[i+1 : j]
}
