package main

func strangePrinter(s string) int {
	res := make([][]int, len(s))
	for i := range res {
		res[i] = make([]int, len(s)+1)
	}

	for i := range s {
		res[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			res[j][i] = res[j][i-1] + 1
			for k := j; k < i; k++ {
				q := res[j][k] + res[k+1][i]
				if s[k] == s[j] && q < res[j][i] {
					res[j][i] = q
				}
			}
			if s[i] == s[j] {
				res[j][i]--
			}
		}
	}

	return res[0][len(s)-1]
}
