package maxvaluecoins

func setMax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func maxValueOfCoins(piles [][]int, k int) int {
	cur := []int{0}
	for _, pile := range piles {
		next := make([]int, k+1)
		for i, c := range cur {
			s := c
			setMax(&next[i], s)
			for j, p := range pile {
				if i+j+1 > k {
					break
				}
				s += p
				setMax(&next[i+j+1], s)
			}
		}
		cur = next
	}
	return cur[k]
}
