package main

func isIsomorphic(s string, t string) bool {
	return canonical(s) == canonical(t)
}

func canonical(s string) string {
	letters := make(map[rune]byte)
	cur := byte('a')
	var res []byte
	for _, c := range s {
		let, ok := letters[c]
		if !ok {
			let = cur
			letters[c] = cur
			cur++
		}
		res = append(res, let)
	}
	return string(res)
}
