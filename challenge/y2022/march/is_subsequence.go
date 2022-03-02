package march

func isSubsequence(s string, t string) bool {
	j := 0
	for _, c := range []byte(s) {
		for j < len(t) && t[j] != c {
			j++
		}
		if j == len(t) {
			return false
		}
	}
	return true
}
