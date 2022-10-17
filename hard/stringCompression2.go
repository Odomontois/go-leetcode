package hard

func encoded(size int) int {
	switch {
	case size == 1:
		return 1
	case size < 10:
		return 2
	case size < 100:
		return 3
	default:
		return 4
	}
}

func getLengthOfOptimalCompression(s string, k int) int {
	var counts [26][]int
	for a := 0; a < 26; a++ {
		counts[a] = make([]int, 1, len(s)+1)
	}
	for i, c := range s {
		for a := 0; a < 26; a++ {
			counts[a] = append(counts[a], counts[a][i])
		}
		counts[c-'a'][i+1]++
	}

	results := make([][]int, len(s))
	for i := range results {
		results[i] = make([]int, k+1)
	}
	var result func(int, int) int
	result = func(start, rem int) int {
		if start == len(s) {
			return 0
		}
		if r := results[start][rem]; r != 0 {
			return r
		}
		best := min(len(s))

		sums := counts[s[start]-'a']
		for i := start + 1; i <= len(s); i++ {
			keep := sums[i] - sums[start]
			del := i - start - keep
			if del > rem {
				break
			}
			best.check(encoded(keep) + result(i, rem-del))
		}
		results[start][rem] = int(best)

		return int(best)
	}

	best := min(len(s))
	for i := 0; i <= k; i++ {
		best.check(result(i, k-i))
	}

	return int(best)
}

type min int

func (best *min) check(a int) {
	if a < int(*best) {
		*best = min(a)
	}
}
