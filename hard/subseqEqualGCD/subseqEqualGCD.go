package subseqequalgcd

func subsequencePairCount(nums []int) int {
	const mod = 1e9 + 7
	all := map[[2]int]int{{0, 0}: 1}
	for _, num := range nums {
		next := map[[2]int]int{}
		for k, v := range all {
			for i := 0; i < 2; i++ {
				k1 := k
				k1[i] = gcd(num, k1[i])
				next[k1] = (next[k1] + v) % mod
			}
		}
		for k, v := range next {
			all[k] = (all[k] + v) % mod
		}
	}
	res := 0
	for k, v := range all {
		if k[0] == k[1] {
			res = (res + v) % mod
		}
	}
	return res - 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
