package robotrepair

import (
	"math"
	"sort"
)

type object struct{ position, limit int64 }

func minimumTotalDistance2(robots []int, factories [][]int) int64 {
	n := len(robots) + len(factories)
	objects := make([]object, 0, n)
	for _, r := range robots {
		objects = append(objects, object{int64(r), -1})
	}
	for _, f := range factories {
		objects = append(objects, object{int64(f[0]), int64(f[1])})
	}

	sort.Slice(objects, func(i, j int) bool { return objects[i].position < objects[j].position })
	objects = append(objects, object{objects[n-1].position + 1, 0})
	roboLimit := make([]int, n+1)
	for i, cur := n, 0; i >= 0; i-- {
		if objects[i].limit < 0 {
			cur++
		}
		roboLimit[i] = cur
	}

	cache := make([]map[int64]int64, n+1)
	for i := range cache {
		cache[i] = make(map[int64]int64)
	}

	var calc func(int, int64) int64
	calc = func(i int, robo int64) int64 {
		if i == n {
			if robo == 0 {
				return 0
			}
			return math.MaxInt64 / 2
		}
		if res, ok := cache[i][robo]; ok {
			return res
		}
		res := int64(math.MaxInt64) / 2
		next := objects[i+1].position - objects[i].position

		if objects[i].limit == -1 {
			if robo < 0 {
				res = calc(i+1, robo+1) - next*(robo+1)
			} else {
				res = calc(i+1, robo+1) + next*(robo+1)
			}
		} else {
			remains := robo - objects[i].limit

			if remains >= 0 {
				res = calc(i+1, remains) + remains*next
			} else {
				maxOrder := int64(-roboLimit[i])
				for order := int64(0); order >= remains && order >= maxOrder; order-- {
					check := calc(i+1, order) - order*next
					if check < res {
						res = check
					}
				}
			}
		}
		cache[i][robo] = res
		return res
	}

	return calc(0, 0)
}
