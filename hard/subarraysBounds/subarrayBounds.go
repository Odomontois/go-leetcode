package subarraysbounds

type event struct{ idx, num int }

func choose(x, y int, cmp func(x, y int) bool) int {
	if cmp(x, y) {
		return x
	}
	return y
}

type MonoStack struct {
	events              []event
	target, start, last int
	cmp                 func(x, y int) bool
}

func (stack *MonoStack) insert(i, x int) {
	es := stack.events
	l := len(es) - 1
	for l >= 0 && stack.cmp(x, es[l].num) {
		l--
	}
	stack.events = append(es[:l+1], event{i, x})
	if x == stack.target {
		stack.last = i + 1
		if l >= 0 {
			stack.start = es[l].idx + 1
		}
	} else if stack.cmp(x, stack.target) {
		stack.last = -1
	}
}

func countSubarrays(nums []int, minK int, maxK int) (res int64) {
	mins := MonoStack{target: minK, cmp: func(x, y int) bool { return x <= y }}
	maxs := MonoStack{target: maxK, cmp: func(x, y int) bool { return x >= y }}
	for i, num := range nums {
		mins.insert(i, num)
		maxs.insert(i, num)
		if mins.last >= 0 && maxs.last >= 0 {
			from := choose(mins.start, maxs.start, maxs.cmp)
			to := choose(mins.last, maxs.last, mins.cmp)
			if to > from {
				res += int64(to - from)
			}
		}
	}
	return res
}
