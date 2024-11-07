package countbalancedperms

import (
	"iter"
	"slices"
)

type distr = [10]int

func countBalancedPermutationsCringe(num string) int {
	total, counts := sumCount[int](num)
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
		res += (perms(t, len(num)/2) * perms(comp, (len(num)+1)/2)) % MOD
	}
	return int(res % MOD)
}

func distribs(num string) iter.Seq[distr] {
	return func(yield func(distr) bool) {
		total, counts := sumCount[int](num)
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
