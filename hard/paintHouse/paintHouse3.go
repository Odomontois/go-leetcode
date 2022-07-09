package painthouse

func minCost(houses []int, cost [][]int, m, _, target int) int {
	type Key struct {
		target, next, i int
	}
	cache := make(map[Key]int)

	var calc func(target, col, i int) int

	calc = func(target, next, i int) int {

		switch {
		case i == -1 && target == 0:
			return 0
		case target > i+1 || target < 0:
			return -1
		}

		key := Key{target, next, i}
		res, ok := cache[key]
		if ok {
			return res
		}

		res = -1

		type Paint struct{ color, cost int }

		var paints []Paint

		if houses[i] == 0 {
			for newCol, paintCost := range cost[i] {
				paints = append(paints, Paint{newCol + 1, paintCost})
			}
		} else {
			paints = append(paints, Paint{houses[i], 0})
		}

		for _, paint := range paints {
			prevTarget := target
			if paint.color != next {
				prevTarget--
			}
			fullCost := calc(prevTarget, paint.color, i-1)
			if fullCost == -1 {
				continue
			}
			fullCost += paint.cost
			if res == -1 || fullCost < res {
				res = fullCost
			}
		}

		cache[key] = res

		return res
	}

	res := calc(target, -1, m-1)

	return res
}
