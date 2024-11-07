package countbalancedperms

const MOD int64 = 1e9 + 7

type distr8 = [10]int8

func countBalancedPermutations(num string) int {
	total, counts := sumCount[int8](num)
	if total%2 == 1 {
		return 0
	}

	var search func(distr8, int, int) int64
	var cache = map[distr8]int{}
	search = func(cur distr8, sum, rest int) int64 {
		if res, ok := cache[cur]; ok {
			return int64(res)
		}
		res := int64(0)

		if sum == 0 && rest == 0 {
			res = perms(cur, (len(num)+1)/2)
		} else if rest > 0 {
			for d := 0; d < 10; d++ {
				if cur[d] > 0 && sum >= d {
					cur[d]--
					res += search(cur, sum-d, rest-1)
					cur[d]++
				}
			}
		} else {
			return 0
		}

		res %= MOD
		cache[cur] = int(res)
		return res
	}
	return int(search(counts, total/2, len(num)/2))
}

var factorials = []int64{1}

func fact(n int) int64 {
	for n >= len(factorials) {
		last := factorials[len(factorials)-1]
		next := (last * int64(len(factorials))) % MOD
		factorials = append(factorials, next)
	}
	return factorials[n]
}

func euclid(a, b int64) (int64, int64) {
	if a == 1 {
		return 1, 0
	}
	x, y := euclid(b, a%b)
	// x b + y r = 1, d b + r = a, r = a - d b
	// x b + y (a - d b) = 1
	// y a + (x - y d) b = 1
	return y, x - y*(a/b)
}

func reci(a int64) int64 {
	x, _ := euclid(a, MOD)
	return (x + MOD) % MOD
}

func comb(n, k int) int64 {
	return (fact(n) * reci((fact(k)*fact(n-k))%MOD)) % MOD
}

func perms[D digit](counts [10]D, sum int) int64 {
	cur := int64(1)
	for _, count := range counts {
		cur = (cur * comb(sum, int(count))) % MOD
		sum -= int(count)
	}
	return cur
}

type digit interface{ int | int8 }

func sumCount[D digit](num string) (total int, counts [10]D) {
	for _, c := range num {
		d := int(c - '0')
		counts[d]++
		total += d
	}
	return
}
