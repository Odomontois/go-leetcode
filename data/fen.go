package fen

type Fen struct {
	size  int
	elems []int
}

func makeFen(size int) Fen {
	return Fen{size, make([]int, size*3)}
}

func (fen Fen) add(index int, amt int) {
	var walk func(int, int, int)
	walk = func(start int, end int, p int) {
		if end <= index || start > index {
			return
		}
		fen.elems[p] += amt
		if end-start == 1 {
			return
		}
		m := (end + start + 1) / 2
		walk(start, m, 2*p+1)
		walk(m, end, 2*p+2)
	}
	walk(0, fen.size, 0)
}

func (fen Fen) calc(from int, until int) int {
	var walk func(int, int, int) int
	walk = func(start int, end int, p int) int {
		if start >= until || end <= from {
			return 0
		}
		if start >= from && end <= until {
			return fen.elems[p]
		}
		m := (end + start + 1) / 2
		return walk(start, m, 2*p+1) + walk(m, end, 2*p+2)
	}
	return walk(0, fen.size, 0)
}
