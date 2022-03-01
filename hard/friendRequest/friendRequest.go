package friendrequest

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	res := make([]bool, 0, len(requests))
	ufs := make([]uf, n)
	for _, req := range requests {
		p1 := ufs[req[0]].getParent()
		p2 := ufs[req[1]].getParent()
		success := true
		for _, rs := range restrictions {
			rp1 := ufs[rs[0]].getParent()
			rp2 := ufs[rs[1]].getParent()
			success = success && (p1 != rp1 && p1 != rp2 || p2 != rp1 && p2 != rp2)
			if !success {
				break
			}
		}
		res = append(res, success)
		if success {
			p1.union(p2)
		}
	}
	return res
}

type uf struct {
	parent *uf
	size   int
}

func (uf *uf) getParent() *uf {
	if uf.parent == nil {
		return uf
	}
	uf.parent = uf.parent.getParent()
	return uf.parent
}

func (p1 *uf) union(p2 *uf) {
	if p1 == p2 {
		return
	}
	if p1.size > p2.size {
		p1.parent = p2
		p2.size += p1.size + 1
	} else {
		p2.parent = p1
		p1.size += p2.size + 1
	}
}
