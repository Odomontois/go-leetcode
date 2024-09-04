package mindaysdisconnect

func minDays(m [][]int) int {
	return mkGrid(m).solve()
}

type point struct{ x, y int }

var noPoint = point{-1, -1}

func (p point) pair() (int, int) {
	return p.x, p.y
}

type grid struct {
	data          [][]int
	m, n          int
	total         int
	level         int
	first, second point
	bridge, cape  bool
}

func mkGrid(data [][]int) (g grid) {
	g = grid{
		data:   data,
		level:  2,
		first:  noPoint,
		second: noPoint,
		bridge: false,
		cape:   false,
		m:      len(data),
		n:      len(data[0]),
	}
	g.forOnes(func(i, j, v int) {
		if g.first == noPoint {
			g.first = point{i, j}
		} else if g.second == noPoint {
			g.second = point{i, j}
		}
		g.total++
		ns := 0
		for _, d := range nds {
			x, y := d.x+i, d.y+j
			if x < 0 || x >= g.m || y < 0 || y >= g.n {
				continue
			}
			ns += data[x][y]
		}
		if ns == 1 {
			g.cape = true
		}
	})
	return g
}

var nds = [4]point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func (g *grid) forOnes(f func(i, j, v int)) {
	for i, r := range g.data {
		for j, v := range r {
			if v > 0 {
				f(i, j, v)
			}
		}
	}
}

func (g *grid) walk(i, j int) int {
	g.data[i][j] = g.level
	g.level++
	level := g.level
	t := -1
	for _, d := range nds {
		d.x += i
		d.y += j
		if d.x < 0 || d.x >= g.m || d.y < 0 || d.y >= g.n {
			continue
		}
		t = g.data[d.x][d.y]
		if t == 0 || t > 1 && t == g.level-2 {
			continue
		}
		if t == 1 {
			t = g.walk(d.pair())
		}
		if t < level {
			level = t
		}
	}
	g.level--
	if t >= 0 && level == g.level+1 {
		g.bridge = true
	}
	return level
}

func (g grid) solve() int {
	if g.first == noPoint {
		return 0
	} else if g.second == noPoint {
		return 1
	}
	g.walk(g.first.pair())
	disconnected := false
	g.forOnes(func(i, j, v int) {
		if v == 1 {
			disconnected = true
		} else {
			g.data[i][j] = 1
		}
	})
	if disconnected {
		return 0
	} else if g.total == 2 {
		return 2
	} else if g.cape {
		return 1
	}
	g.walk(g.second.pair())
	if g.bridge {
		return 1
	}
	return 2
}
