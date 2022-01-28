package march

func wordSubsets(A []string, B []string) (res []string) {
	total, req := [26]int16{}, 0
	for _, b := range B {
		current := [26]int16{}
		for _, c := range b {
			x := c - 'a'
			current[x]++
			if total[x] < current[x] {
				total[x]++
			}
		}
	}
	for _, t := range total {
		if t != 0 {
			req++
		}
	}
	for _, a := range A {
		current, gain := [26]int16{}, 0
		for _, c := range a {
			x := c - 'a'
			if current[x] < 0 {
				continue
			}
			current[x]++
			if total[x] > 0 && current[x] >= total[x] {
				current[x] = -1
				gain++
			}
			if gain == req {
				res = append(res, a)
				break
			}
		}
	}
	return
}