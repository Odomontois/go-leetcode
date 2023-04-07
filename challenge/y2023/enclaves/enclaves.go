package enclaves

func numEnclaves(grid [][]int) int {
	n, m := len(grid)-1, len(grid[0])-1
	var descend func(i, j int) (int, bool)

	descend = func(i, j int) (int, bool) {
		if grid[i][j] == 0 {
			return 0, true
		}
		if i == 0 || i == n || j == 0 || j == m {
			return -1, false
		}
		grid[i][j] = 0
		res, closed := 1, true
		for _, p := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			d, c := descend(i+p[0], j+p[1])
			res += d
			closed = closed && c
		}

		return res, closed
	}

	res := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				if d, closed := descend(i, j); closed {
					res += d
				}
			}
		}
	}
	return res
}
