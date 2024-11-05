package countbalancedperms

func countBalancedPermutationsOld(num string) int {
	const mod int = 1e9 + 7
	const mod64 = int64(mod)

	total, counts := 0, [10]int8{}
	for _, c := range num {
		d := int(c - '0')
		counts[d]++
		total += d
	}
	type cache map[[10]int8]int
	gen := func(prev cache) cache {
		res := make(cache)
		for k, v := range prev {
			for d := 0; d < 10; d++ {
				if k[d] < counts[d] {
					k[d]++
					res[k] = (res[k] + v) % mod
					k[d]--
				}
			}
		}
		return res
	}
	cur := cache{{}: 1}
	for i := 0; i < len(num)/2; i++ {
		cur = gen(cur)
	}
	next := cur
	if len(num)%2 == 1 {
		next = gen(cur)
	}
	result := 0
	for k, v := range next {
		s, comp := 0, counts
		for d, c := range k {
			s += d * int(c)
			comp[d] -= c
		}
		if s*2 == total {
			result = (result + int(int64(v)*int64(cur[comp])%mod64)) % mod
		}
	}
	return result
}
