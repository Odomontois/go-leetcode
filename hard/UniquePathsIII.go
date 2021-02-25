package hard

func UniquePathsIII(grid [][]int) int {
	return uniquePathsIII(grid)
}

func uniquePathsIII(grid [][]int) int {
	seen := make([][]bool, len(grid))
	n, m := len(grid), len(grid[0])
	rem := 0
	for i, row := range grid {
		seen[i] = make([]bool, len(grid[i]))
		for _, c := range row {
			if c == 0 || c == 1 {
				rem++
			}
		}
	}
	var walk func(int, int) int
	walk = func(i int, j int) int {
		if seen[i][j] || grid[i][j] == -1 || grid[i][j] == 2 && rem > 0 {
			return 0
		}
		if grid[i][j] == 2 && rem == 0 {
			return 1
		}
		seen[i][j] = true
		rem--
		res := 0
		for _, p := range [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
			x, y := p[0], p[1]
			if x >= 0 && x < n && y >= 0 && y < m {
				res += walk(x, y)
			}
		}
		seen[i][j] = false
		rem++
		return res
	}
	for i, row := range grid {
		for j, c := range row {
			if c == 1 {
				return walk(i, j)
			}
		}
	}
	return 0
}
