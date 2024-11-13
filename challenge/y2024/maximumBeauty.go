package y2024

import "slices"

func maximumBeauty(items [][]int, queries []int) []int {
	type query struct{ price, idx int }
	type item struct{ price, beauty int }
	is := make([]item, len(items))
	for i, it := range items {
		is[i] = item{it[0], it[1]}
	}
	slices.SortFunc(is, func(a, b item) int { return a.price - b.price })
	qs := make([]query, len(queries))
	for i, q := range queries {
		qs[i].idx = i
		qs[i].price = q
	}
	slices.SortFunc(qs, func(a, b query) int { return a.price - b.price })
	res := make([]int, len(queries))
	var cur, maxBeauty int
	for _, item := range is {
		for cur < len(qs) && qs[cur].price < item.price {
			res[qs[cur].idx] = maxBeauty
			cur++
		}
		if item.beauty > maxBeauty {
			maxBeauty = item.beauty
		}
	}
	for _, q := range qs[cur:] {
		res[q.idx] = maxBeauty
	}
	return res
}
