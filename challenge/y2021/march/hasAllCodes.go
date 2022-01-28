package march

func hasAllCodes(s string, k int) bool {
	codes, remain, code, mask := make([]bool, 1<<k), 1<<k, 0, (1<<k)-1
	for i, c := range s {
		code = mask&(code<<1) | int(c-'0')
		if i < k-1 {
			continue
		}
		if !codes[code] {
			remain--
		}
		codes[code] = true
	}
	return remain == 0
}
