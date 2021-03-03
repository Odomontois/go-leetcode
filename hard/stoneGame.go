package hard

import "sort"

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
func StoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n <= 1 {
		return 0
	}
	games, sums, cur := make([][]int, n), make([]int, n+1), 0
	for i := range games {
		games[i] = make([]int, n-i-1)
		cur += stoneValue[i]
		sums[i+1] = cur
	}
	var gain func(int, int, int) int
	game := func(i int, k int) int {
		if k == i {
			return 0
		}
		if games[i][k-i-1] == 0 {
			best := 0
			j1 := max(sort.SearchInts(sums, (sums[i]+sums[k+1])/2)-2, i)
			for j := j1; j <= j1+1 && j < k; j++ {
				best = max(gain(i, j, k), best)
			}
			games[i][k-i-1] = best
		}
		return games[i][k-i-1]
	}
	gain = func(i int, j int, k int) int {
		left, right := game(i, j), game(j+1, k)
		lsum, rsum := sums[j+1]-sums[i], sums[k+1]-sums[j+1]
		if lsum > rsum || lsum == rsum && right > left {
			return rsum + right
		}
		return lsum + left
	}
	return game(0, n-1)
}
