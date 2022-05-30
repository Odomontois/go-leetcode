package minobstaclerem

// https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/

func minimumObstacles(grid [][]int) int {
	var cur, next [][2]int
	n, m := len(grid), len(grid[0])
	steps := 0
	cur = append(cur, [2]int{0, 0})
	for {
		for len(cur) > 0 {
			x := cur[0]
			cur = cur[1:]
			i, j := x[0], x[1]
			if i == n-1 && j == m-1 {
				return steps
			}
			if grid[i][j] < 0 {
				continue
			}
			grid[i][j] = -1
			for _, y := range [...][2]int{{i - 1, j}, {i + 1, j}, {i, j + 1}, {i, j - 1}} {
				i, j := y[0], y[1]
				if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] < 0 {
					continue
				}
				if grid[i][j] == 1 {
					next = append(next, y)
				} else {
					cur = append(cur, y)
				}
			}
		}
		cur, next = next, nil
		steps++
	}
}
