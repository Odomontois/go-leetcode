package hard

func areConnected(n int, threshold int, queries [][]int) (result []bool) {
	ds := make([]disjoint, n+1)
	result = make([]bool, len(queries))
	for i := range ds {
		ds[i].size = 1
	}
	for k := threshold + 1; 2*k <= n; k++ {
		for j := 2 * k; j <= n; j += k {
			ds[k].Join(&ds[j])
		}
	}

	for i, q := range queries {
		result[i] = ds[q[0]].GetParent() == ds[q[1]].GetParent()
	}
	return
}

type disjoint struct {
	parent *disjoint
	size   int
}

func (ds *disjoint) GetParent() *disjoint {
	if ds.parent == nil {
		return ds
	}
	p := ds.parent.GetParent()
	ds.parent = p
	return p
}

func (ds *disjoint) Join(other *disjoint) {
	lp, rp := ds.GetParent(), other.GetParent()
	if lp == rp {
		return
	}
	if lp.size > rp.size {
		rp.parent = lp
		lp.size += rp.size
	} else {
		lp.parent = rp
		rp.size += lp.size
	}
}
