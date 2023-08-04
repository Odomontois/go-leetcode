package wordbreak

func wordBreak(s string, wordDict []string) bool {
	res := make([]bool, len(s)+1)
	res[0] = true
	for i := range s {
		for _, word := range wordDict {
			w := len(word)
			start := i + 1 - w
			if start >= 0 && res[start] && s[start:i+1] == word {
				res[i+1] = true
				break
			}
		}
	}
	return res[len(s)]
}
