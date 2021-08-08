package main

func minCut(s string) int {
	best := make([]int, len(s)+1)
	for i := range best {
		best[i] = i
	}
	for k := 0; k < 2*len(s)+1; k++ {
		i, j := k/2, (k+1)/2
		for i >= 0 && j < len(s) && s[i] == s[j] {
			if best[j+1] > best[i]+1 {
				best[j+1] = best[i] + 1
			}
			i--
			j++
		}
	}
	return best[len(s)] - 1
}
