package trionic2

import "math"

type stage struct {
	sum    int64
	active bool
}

func (stage *stage) add(init int64, value int, cond bool, start bool) {
	if !cond {
		stage.active = false
		return
	}
	if start {
		stage.active = true
		stage.sum = init + int64(value)
	} else if stage.active {
		stage.sum += int64(value)
	}
}

func maxSumTrionic(nums []int) int64 {
	var inc, dec, trio stage
	var best int64 = math.MinInt64
	for i, cur := range nums[1:] {
		prev := nums[i]
		p64 := int64(prev)

		trio.add(dec.sum, cur, cur > prev, dec.active)
		dec.add(inc.sum, cur, cur < prev, inc.active)
		inc.add(p64, cur, cur > prev, !inc.active || p64 > inc.sum)

		if trio.active && trio.sum > best {
			best = trio.sum
		}
	}
	return best
}
