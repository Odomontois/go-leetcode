package wildcard

func isMatch(s string, p string) bool {
	res := make([][]bool, len(p)+1)
	for i := range res {
		res[i] = make([]bool, len(s)+1)
	}
	res[len(p)][len(s)] = true
	for k := len(p) - 1; k >= 0; k-- {
		if p[k] == '*' {
			res[k][len(s)] = res[k+1][len(s)]
		}
		for j := len(s) - 1; j >= 0; j-- {
			r := false
			switch p[k] {
			case '?', s[j]:
				r = res[k+1][j+1]
			case '*':
				r = res[k+1][j] || res[k][j+1]
			}
			res[k][j] = r
		}
	}
	return res[0][0]
}
