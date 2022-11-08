package robotrepair

import (
	"math"
	"sort"
)

type robot int64

func (r robot) position() int64 { return int64(r) }

type factory struct {
	pos, limit int64
}

func (f factory) position() int64 { return f.pos }

type position interface {
	position() int64
}

type end int64

func (e end) position() int64 { return int64(e) }

func minimumTotalDistance(robots []int, factories [][]int) int64 {
	n := len(robots) + len(factories)
	objects := make([]position, 0, n)
	for _, r := range robots {
		objects = append(objects, robot(r))
	}
	for _, f := range factories {
		objects = append(objects, factory{int64(f[0]), int64(f[1])})
	}

	sort.Slice(objects, func(i, j int) bool { return objects[i].position() < objects[j].position() })
	objects = append(objects, end(objects[n-1].position()+1))
	roboLimit := make([]int, n+1)
	for i, cur := n, 0; i >= 0; i-- {
		switch objects[i].(type) {
		case robot:
			cur++
		}
		roboLimit[i] = cur
	}

	cache := make([]map[int64]int64, n+1)
	for i := range cache {
		cache[i] = make(map[int64]int64)
	}

	pos := func(i int) int64 { return objects[i].position() }

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
		next := pos(i+1) - pos(i)

		switch objects[i].(type) {
		case robot:
			if robo < 0 {
				res = calc(i+1, robo+1) - next*(robo+1)
			} else {
				res = calc(i+1, robo+1) + next*(robo+1)
			}
		case factory:
			limit := objects[i].(factory).limit
			remains := robo - limit

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
