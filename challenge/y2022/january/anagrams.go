package main

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s, p string) []int {
	var pat, window [26]int
	matched, n := 26, len(p)
	for _, c := range letterNums(p) {
		if pat[c] == 0 {
			matched--
		}
		pat[c]++
	}

	var acc []int
	sb := letterNums(s)

	mutate := func(c byte, delta int) {
		if pat[c] == window[c] {
			matched--
		}
		window[c] += delta
		if pat[c] == window[c] {
			matched++
		}
	}

	for i, c := range sb {
		if i >= n {
			mutate(sb[i-n], -1)
		}
		mutate(c, 1)
		if matched == 26 {
			acc = append(acc, i-n+1)
		}
	}

	return acc
}

func letterNums(s string) []byte {
	const A = byte('a')
	res := []byte(s)
	for i := range res {
		res[i] -= A
	}
	return res
}
