package similararrays

import "sort"

func abs(x int) int64 {
	if x < 0 {
		return -int64(x)
	}
	return int64(x)
}
func makeSimilar(nums []int, target []int) (res int64) {
	counter := make(map[int]int)
	for _, x := range nums {
		counter[x]++
	}
	for _, x := range target {
		counter[x]--
	}
	type accum struct{ num, count int }
	acc := make([]accum, 0, len(counter))
	for k, v := range counter {
		acc = append(acc, accum{k, v})
	}
	sort.Slice(acc, func(i, j int) bool { return acc[i].num < acc[j].num })
	x := acc[0].num
	cur, next := accum{x, 0}, accum{x + 1, 0}
	i := 0
	for i < len(acc) {
		if cur.num < acc[i].num {
			res += abs(next.count)
			cur.num += 2
			cur, next = next, cur
		} else {
			cur.count += acc[i].count
			i++
		}
	}

	return res / 2
}
