package friendrequest

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	res := make([]bool, 0, len(requests))
	uf := make(uf, n)
	for _, req := range requests {
		p1 := uf.getParent(req[0])
		p2 := uf.getParent(req[1])
		success := true
		for _, rs := range restrictions {
			rp1 := uf.getParent(rs[0])
			rp2 := uf.getParent(rs[1])
			success = success && (p1 != rp1 || p2 != rp2) && (p2 != rp1 || p1 != rp2)
			if !success {
				break
			}
		}
		res = append(res, success)
		if success {
			uf.union(p1, p2)
		}
	}
	return res
}

type uf []int

func (uf uf) getParent(i int) int {
	if uf[i] <= 0 {
		return i
	}
	t := uf.getParent(uf[i] - 1)
	uf[i] = t + 1
	return t
}

func (uf uf) union(p1, p2 int) {
	if p1 == p2 {
		return
	}
	size := uf[p1] + uf[p2] - 1
	if uf[p1] < uf[p2] {
		uf[p1], uf[p2] = size, p1+1
	} else {
		uf[p1], uf[p2] = p2+1, size
	}
}
