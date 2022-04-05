package hard

func sumScores(s string) int64 {
	z := make([]int, len(s))
	l, r, res, n := 0, 0, int64(0), len(s)
	for i := 1; i < n; i++ {
		p := 0
		if i <= r {
			p = z[i-l]
			if r-i+1 < p {
				p = r - i + 1
			}
		}
		for i+p < n && s[p] == s[i+p] {
			p++
		}
		if i+p-1 > r {
			l = i
			r = i + p - 1
		}
		z[i] = p
		res += int64(p)
	}
	return res + int64(n)
}
