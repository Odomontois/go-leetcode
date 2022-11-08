package makearrayequal

func sign(x int) int64 {
	if x < 0 {
		return -1
	}
	return 1
}
func abs(x int) int64 {
	if x < 0 {
		return int64(-x)
	}
	return int64(x)
}
func minCost(nums []int, costs []int) int64 {
	cost := func(x int, f func(x int) int64) (res int64) {
		for i, n := range nums {
			res += f(x-n) * int64(costs[i])
		}
		return res
	}

	max, min := nums[0], nums[0]
	for _, x := range nums {
		if x > max {
			max = x
		} else if x < min {
			min = x
		}
	}
	for max-min > 1 {
		m := (max + min) / 2
		if cost(m, sign) < 0 {
			min = m
		} else {
			max = m
		}
	}
	minCost, maxCost := cost(min, abs), cost(max, sign)
	if minCost < maxCost {
		return minCost
	}
	return maxCost
}
