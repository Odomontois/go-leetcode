package secondgreater

type item struct{ idx, max int }

func (me item) choose(that item) item {
	if me.max > that.max {
		return me
	}
	return that
}

type pos struct{ idx, from, to int }

func (p pos) split() (pos, pos) {
	m, l, r := (p.from+p.to+1)/2, p.idx*2+1, p.idx*2+2
	return pos{l, p.from, m}, pos{r, m, p.to}
}

func (p pos) last() bool {
	return p.to-p.from == 1
}

func secondGreaterElement(nums []int) []int {
	maxTree := make([]item, len(nums)*4)
	var build func(pos)
	init := pos{0, 0, len(nums)}
	none := item{-1, -1}
	build = func(p pos) {
		if p.last() {
			maxTree[p.idx].idx = p.from
			maxTree[p.idx].max = nums[p.from]
			return
		}
		l, r := p.split()
		build(l)
		build(r)
		maxTree[p.idx] = maxTree[l.idx].choose(maxTree[r.idx])
	}
	build(init)
	search := func(start, min int) item {
		var iter func(pos) item
		iter = func(p pos) item {
			if maxTree[p.idx].max <= min || p.to <= start {
				return none
			}
			if p.last() {
				return maxTree[p.idx]
			}
			l, r := p.split()
			res := iter(l)
			if res != none {
				return res
			}
			return iter(r)
		}
		return iter(init)
	}
	res := make([]int, len(nums))
	for i, x := range nums {
		res[i] = -1
		first := search(i+1, x)
		if first == none {
			continue
		}
		second := search(first.idx+1, x)
		if second == none {
			res[i] = second.max
		}

	}
	return res
}
