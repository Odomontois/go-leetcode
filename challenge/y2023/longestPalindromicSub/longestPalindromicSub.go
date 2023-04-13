package longestpalindromicsub

func longestPalindromeSubseq(s string) int {
	cache := make(map[[2]int]int)
	var calc func(int, int) int
	calc = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		k := [2]int{i, j}
		if v, ok := cache[k]; ok {
			return v
		}
		best := 0
		if s[i] == s[j] {
			best = calc(i+1, j-1) + 2
		} else {
			best = calc(i+1, j)
			r := calc(i, j-1)
			if best < r {
				best = r
			}
		}
		cache[k] = best
		return best
	}
	return calc(0, len(s)-1)
}
