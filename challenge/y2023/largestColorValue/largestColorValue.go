package largestcolorvalue

func largestPathValue(colors string, edges [][]int) int {
	graph := make([]struct {
		out    []int
		in     []int
		outCnt int
		max    [26]int
	}, len(colors))
	for _, e := range edges {
		out := &graph[e[0]].out
		*out = append(*out, e[1])
		in := &graph[e[1]].in
		*in = append(*in, e[0])
		graph[e[0]].outCnt++
	}
	var q []int
	for i, v := range graph {
		if v.outCnt == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		mijn := &graph[i]
		for _, j := range mijn.in {
			graph[j].outCnt--
			if graph[j].outCnt == 0 {
				q = append(q, j)
			}
		}
		for _, j := range mijn.out {
			zijn := &graph[j]
			for k := 0; k < 26; k++ {
				if zijn.max[k] > mijn.max[k] {
					mijn.max[k] = zijn.max[k]
				}
			}
		}
		mijn.max[colors[i]-'a']++
	}

	best := 0
	for _, v := range graph {
		if v.outCnt > 0 {
			return -1
		}
		for _, m := range v.max {
			if m > best {
				best = m
			}
		}
	}

	return best
}
