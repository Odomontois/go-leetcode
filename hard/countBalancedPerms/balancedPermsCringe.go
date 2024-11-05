package countbalancedperms

import (
	"iter"
	"slices"
)

const MOD int64 = 1e9 + 7

type distr = [10]int

func countBalancedPermutationsCringe(num string) int {
	total, counts := sumCount(num)
	if total%2 == 1 {
		return 0
	}

	target := slices.AppendSeq([]distr(nil), distribs(num))
	res := int64(0)
	for _, t := range target {
		var comp [10]int
		for d, c := range t {
			comp[d] = counts[d] - c
		}
		res += (perms(t[:], len(num)/2) * perms(comp[:], (len(num)+1)/2)) % MOD
	}
	return int(res % MOD)
}

func sumCount(num string) (total int, counts distr) {
	for _, c := range num {
		d := int(c - '0')
		counts[d]++
		total += d
	}
	return
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

func perms(counts []int, sum int) int64 {
	cur := int64(1)
	for _, count := range counts {
		cur = (cur * comb(sum, count)) % MOD
		sum -= count
	}
	return cur
}

func distribs(num string) iter.Seq[distr] {
	return func(yield func(distr) bool) {
		total, counts := sumCount(num)
		csums := counts
		digSums := [10]int{}
		for d := 1; d < 10; d++ {
			csums[d] += csums[d-1]
			digSums[d] = digSums[d-1] + d*counts[d]
		}
		var recur func(int, int, int) bool
		var cur distr
		recur = func(d int, s int, rest int) bool {
			if d == -1 && s == 0 && rest == 0 {
				return yield(cur)
			} else if d == -1 || csums[d] < rest || digSums[d] < s {
				return true
			}
			for cur[d] = 0; cur[d]*d <= s && cur[d] <= counts[d] && cur[d] <= rest; cur[d]++ {
				if !recur(d-1, s-cur[d]*d, rest-cur[d]) {
					return false
				}
			}
			return true
		}
		recur(9, total/2, len(num)/2)
	}
}
