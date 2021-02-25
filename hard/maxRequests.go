package hard

func maximumRequests(n int, requests [][]int) int {
	best := 0
	for f := 0; f < 1<<len(requests); f++ {
		b := make([]int, n)
		nz := 0
		k := 0
		add := func(i int, d int) {
			if b[i] == 0 {
				nz++
			}
			b[i] += d
			if b[i] == 0 {
				nz--
			}
		}
		for i, req := range requests {
			if f&(1<<i) == 0 {
				continue
			}
			k++
			add(req[0], 1)
			add(req[1], -1)
		}

		if nz == 0 && k > best {
			best = k
		}
	}
	return best
}

// func main() {
// 	fmt.Println(maximumRequests(6, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}))
// }
