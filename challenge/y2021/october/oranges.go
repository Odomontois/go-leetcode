package october

func orangesRotting(grid [][]int) int {
	fresh := 0
	var rotten [][2]int
	for i, line := range grid {
		for j, cell := range line {
			switch cell {
			case 1:
				fresh++
			case 2:
				rotten = append(rotten, [2]int{i, j})
			}
		}
	}
	n, m := len(grid), len(grid[0])
	time := 0
	for ; len(rotten) > 0 && fresh > 0; time++ {
		var next [][2]int
		for _, p := range rotten {
			x, y := p[0], p[1]
			for _, q := range [...][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
				a, b := q[0], q[1]
				if a >= 0 && a < n && b >= 0 && b < m && grid[a][b] == 1 {
					fresh--
					grid[a][b] = 2
					next = append(next, [2]int{a, b})
				}
			}
		}
		rotten = next
	}
	if fresh > 0 {
		time = -1
	}
	return time
}
