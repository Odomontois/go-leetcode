package hard

func minDifficulty(jobDifficulty []int, d int) int {
	results := make([][]int, d+1)
	for i := range results {
		results[i] = make([]int, len(jobDifficulty)+1)
	}
	for i := 0; i < len(jobDifficulty); i++ {
		for k := 0; k < d; k++ {
			res := results[k][i]
			if res == 0 && k != 0 {
				continue
			}
			max := jobDifficulty[i]
			for j := i + 1; j <= len(jobDifficulty); j++ {
				if max < jobDifficulty[j-1] {
					max = jobDifficulty[j-1]
				}
				res1 := &results[k+1][j]
				if *res1 == 0 || *res1 > res+max+1 {
					*res1 = res + max + 1
				}
			}
		}

	}
	res := results[d][len(jobDifficulty)]
	if res == 0 {
		return -1
	}
	return res - d
}
