package paintgrid3colors

import "slices"

func genPatterns(i, mask, prev int, res *[]int) {
	for c := 0; c < 3; c++ {
		if c == prev {
			continue
		}
		nextMask := mask | (1 << (i*3 + c))
		if i == 0 {
			*res = append(*res, nextMask)
		} else {
			genPatterns(i-1, nextMask, c, res)
		}
	}
}

const MOD = 1000000007

func colorTheGrid(m int, n int) (res int) {
	var patterns []int
	genPatterns(m-1, 0, -1, &patterns)
	pn := len(patterns)
	edges := make([][]int, pn)
	for i := range edges {
		for j := range patterns {
			if patterns[i]&patterns[j] == 0 {
				edges[i] = append(edges[i], j)
			}
		}
	}
	cur := slices.Repeat([]int{1}, pn)
	next := make([]int, pn)
	for i := 1; i < n; i++ {
		for i, r := range edges {
			for _, j := range r {
				next[j] = (next[j] + cur[i]) % MOD
			}
		}
		cur, next = next, cur
		clear(next)
	}
	for _, v := range cur {
		res = (res + v) % MOD
	}
	return res
}
