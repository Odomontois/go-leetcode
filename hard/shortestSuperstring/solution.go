package shortestSuperstring

import "strings"

type Result struct {
	prev int8
	len  uint8
}

func shortestSuperstring(words []string) string {
	commons := make([][]int, len(words))
	for i := range commons {
		commons[i] = make([]int, len(words))
		for j := range commons[i] {
			commons[i][j] = len(words[j]) - common(words[i], words[j])
		}
	}
	best := make([][]Result, len(words))
	n := 1 << len(words)
	for i, w := range words {
		best[i] = make([]Result, n)
		best[i][1<<i] = Result{prev: -1, len: uint8(len(w))}
	}

	update := func(v *Result, k int, prev int) {
		if v.len == 0 || int(v.len) > k {
			v.len = uint8(k)
			v.prev = int8(prev)
		}
	}

	for u := 1; u < n; u++ {
		for i := range words {
			if u&(1<<i) == 0 {
				continue
			}
			cur := best[i][u]
			for j := range words {
				update(&best[j][u|(1<<j)], int(cur.len)+commons[i][j], i)
			}
		}
	}
	var final Result
	for i := range words {
		update(&final, int(best[i][n-1].len), i)
	}
	prev, result, cut, flags := final.prev, "", 0, n-1
	for prev >= 0 {
		p := prev
		w := words[p]
		result = w[:len(w)-cut] + result
		prev = best[p][flags].prev
		flags ^= 1 << p
		if prev >= 0 {
			cut = len(w) - commons[prev][p]
		}
	}

	return result
}

func common(s1, s2 string) int {
	for i := len(s2); i >= 0; i-- {
		if strings.HasSuffix(s1, s2[:i]) {
			return i
		}
	}
	return 0
}
