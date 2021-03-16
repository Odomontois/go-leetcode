package bestRotation

func bestRotation(a []int) int {
	var iter func(int, int) Intervals
	iter = func(from, to int) Intervals {
		if to == from+1 {
			return Single(from, a[from], len(a))
		}
		m := (to + from + 1) / 2
		return iter(from, m).Plus(iter(m, to))
	}
	var best Interval
	for _, p := range iter(0, len(a)) {
		if best.points < p.points {
			best = p
		}
	}
	return int(best.shift)
}

type Interval struct {
	shift, points uint16
}

type Intervals []Interval

func Of(x, n int) int {
	return (x + n) % n
}

func Single(i, a, n int) Intervals {
	x, y := Of(i+1, n), Of(i-a, n)
	if y+1 == n {
		return Intervals{{uint16(x), 1}}
	}
	if x <= y {
		return Intervals{{uint16(x), 1}, {uint16(y) + 1, 0}}
	}
	return Intervals{{0, 1}, {uint16(y) + 1, 0}, {uint16(x), 1}}
}

func (in Intervals) Plus(other Intervals) (res Intervals) {
	cura, curb := uint16(0), uint16(0)
	for len(in) > 0 || len(other) > 0 {
		var next uint16
		if len(in) == 0 || (len(other) > 0 && in[0].shift > other[0].shift) {
			next = other[0].shift
			curb, other = other[0].points, other[1:]
		} else if len(other) == 0 || in[0].shift < other[0].shift {
			next = in[0].shift
			cura, in = in[0].points, in[1:]
		} else {
			next = in[0].shift
			cura, in = in[0].points, in[1:]
			curb, other = other[0].points, other[1:]
		}
		res = append(res, Interval{next, cura + curb})
	}
	return
}
