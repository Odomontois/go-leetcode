package march

import "sort"

func coinChange1(coins []int, amount int) int {
	sort.Ints(coins)
	var calc func(int) int
	cache := make(map[int]int)

	compute := func(amt int, res chan int) {
		best := -1
		for i := 0; i < len(coins) && coins[i] <= amt; i++ {
			cur := calc(amt - coins[i])
			if cur != -1 && (best == -1 || best > cur) {
				best = cur + 1
			}
		}
		res <- best
	}

	calc = func(amt int) int {
		if amt == 0 {
			return 0
		}
		cached, found := cache[amt]
		if found {
			return cached
		}

		out := make(chan int)
		go compute(amt, out)
		res := <-out
		cache[amt] = res
		return res
	}

	return calc(amount)
}

func coinChange2(coins []int, amount int) int {
	sort.Ints(coins)
	var calc func(int) int
	cache := make(map[int]int)

	calc = func(amt int) int {
		if amt == 0 {
			return 0
		}
		cached, found := cache[amt]
		if found {
			return cached
		}

		best := -1
		for i := 0; i < len(coins) && coins[i] <= amt; i++ {
			cur := calc(amt - coins[i])
			if cur != -1 && (best == -1 || best > cur) {
				best = cur + 1
			}
		}
		cache[amt] = best
		return best
	}

	return calc(amount)
}

func coinChange3(coins []int, amount int) int {
	sort.Ints(coins)
	var calc func(int) int
	cache := make([]int, amount+1)

	calc = func(amt int) int {
		if amt == 0 {
			return 0
		}
		cached := cache[amt]
		if cached > 0 {
			return cached
		}

		best := -1
		for i := 0; i < len(coins) && coins[i] <= amt; i++ {
			cur := calc(amt - coins[i])
			if cur != -1 && (best == -1 || best > cur) {
				best = cur + 1
			}
		}
		cache[amt] = best
		return best
	}

	return calc(amount)
}

func coinChange4(coins []int, amount int) int {
	sort.Ints(coins)
	var calc func(int) int
	cache := make([]int, amount+1)

	compute := func(amt int, res chan int) {
		best := -1
		for i := 0; i < len(coins) && coins[i] <= amt; i++ {
			cur := calc(amt - coins[i])
			if cur != -1 && (best == -1 || best > cur) {
				best = cur + 1
			}
		}
		res <- best
	}

	calc = func(amt int) int {
		if amt == 0 {
			return 0
		}
		cached := cache[amt]
		if cached != 0 {
			return cached
		}

		out := make(chan int)
		go compute(amt, out)
		res := <-out
		cache[amt] = res
		return res
	}

	return calc(amount)
}

func coinChange5(coins []int, amount int) int {
	sort.Ints(coins)
	res := make([]int, amount+1)
	for x, c := range res {
		if x != 0 && c == 0 {
			continue
		}
		for i := 0; i <= len(coins) && coins[i]+x <= amount; i++ {
			y := coins[i] + x
			if res[y] == 0 || res[y] > c+1 {
				res[y] = c + 1
			}
		}
	}
	if res[amount] == 0 && amount != 0 {
		return -1
	}

	return res[amount]
}
