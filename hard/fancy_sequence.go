package hard

type Fancy struct {
	cap  int
	top  []Lin
	vals []Mnum
}

func Constructor() Fancy {
	return Fancy{1, []Lin{}, []Mnum{}}
}

func (fancy *Fancy) extend() {
	fancy.cap = fancy.cap * 2

	top := make([]Lin, len(fancy.top)*2+1)
	for p := range top {
		top[p].mul = 1
	}
	for p, q := 1, 0; p < len(top); p, q = p*2+1, q*2+1 {
		for i := 0; i <= q; i++ {
			top[p+i] = fancy.top[q+i]
		}
	}
	fancy.top = top
}

func (fancy *Fancy) Append(val int) {
	fancy.vals = append(fancy.vals, Mnum(val))
	if len(fancy.vals) > fancy.cap {
		fancy.extend()
	}
}

func (fancy *Fancy) recurMod(lin Lin, start int, end int) {
	var rec func(Lin, int, int, int)
	rec = func(lin Lin, p int, from int, until int) {
		if end <= from || start >= until {
			return
		}
		if until-from == 1 {
			fancy.vals[from] = lin.Apply(fancy.vals[from])
			return
		}
		if from >= start && until <= end {
			fancy.top[p] = fancy.top[p].Combine(lin)
			return
		}
		next := fancy.top[p].Combine(lin)
		fancy.top[p] = Lin{mul: 1, add: 0}
		m := (from + until + 1) / 2
		rec(next, p*2+1, from, m)
		rec(next, p*2+2, m, until)
	}
	rec(lin, 0, 0, fancy.cap)
}

func (fancy *Fancy) AddAll(inc int) {
	fancy.recurMod(Lin{add: Mnum(inc), mul: 1}, 0, len(fancy.vals))
}

func (fancy *Fancy) MultAll(m int) {
	fancy.recurMod(Lin{add: 0, mul: Mnum(m)}, 0, len(fancy.vals))
}

func (fancy *Fancy) GetIndex(idx int) int {
	if idx >= len(fancy.vals) {
		return -1
	}
	var rec func(int, int, int) Mnum
	rec = func(p int, from int, until int) Mnum {
		if until-from == 1 {
			return fancy.vals[from]
		}
		m := (from + until + 1) / 2
		var res Mnum
		if idx < m {
			res = rec(p*2+1, from, m)
		} else {
			res = rec(p*2+2, m, until)
		}
		return fancy.top[p].Apply(res)
	}
	return int(rec(0, 0, fancy.cap))
}

type Lin struct {
	mul Mnum
	add Mnum
}

func (lin Lin) Combine(other Lin) Lin {
	return Lin{
		mul: lin.mul.Mul(other.mul),
		add: lin.add.Mul(other.mul).Plus(other.add),
	}
}

func (lin Lin) Apply(num Mnum) Mnum {
	return lin.mul.Mul(num).Plus(lin.add)
}
