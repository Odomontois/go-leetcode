package hard

func validArrangement(pairs [][]int) [][]int {
	vs := make(map[int]struct {
		ins  int
		outs []int
	})
	open := make(map[int]int)
	for _, p := range pairs {
		v := vs[p[1]]
		v.ins++
		vs[p[1]] = v
		v = vs[p[0]]
		v.outs = append(v.outs, p[1])
		vs[p[0]] = v
	}
	start, end := -1, -1
	for v, p := range vs {
		switch {
		case p.ins < len(p.outs):
			start = v
		case len(p.outs) == 0:
			end = v
		case len(p.outs) < p.ins:
			end = v
		case start == -1 && end == -1:
			start = v
			end = v
		}
	}
	chain := make([][2]int, 0, len(pairs)+1)
	chain = append(chain, [2]int{start, -1})
	link := func(cur, end, prev int) {
		for v := vs[cur]; len(v.outs) > 0; v = vs[cur] {
			next := v.outs[len(v.outs)-1]
			v.outs = v.outs[:len(v.outs)-1]
			i := len(chain)
			chain = append(chain, [2]int{next, end})
			chain[prev][1] = i
			vs[cur] = v
			if len(v.outs) > 0 {
				open[cur] = prev
			} else {
				delete(open, cur)
			}
			prev = i
			cur = next
		}
	}
	link(start, -1, 0)
	for len(open) > 0 {
		for first, c := range open {
			x := chain[c]
			link(first, x[1], c)
			break
		}
	}
	var res [][]int
	prev := start
	for i := chain[0][1]; i != -1; i = chain[i][1] {
		res = append(res, []int{prev, chain[i][0]})
		prev = chain[i][0]
	}
	return res
}
