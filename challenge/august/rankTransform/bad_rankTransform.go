package main

// wrong, comparing only neighbours
func badMatrixRankTransform(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	ranks := make([][]int, n)
	for i := range ranks {
		ranks[i] = make([]int, m)
	}

	var level [][2]int
	for i := range ranks {
		for j := range ranks[i] {
			for _, xy := range [...][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
				x, y := xy[0], xy[1]
				if x >= 0 && y >= 0 && x < n && y < m && matrix[x][y] > matrix[i][j] {
					ranks[x][y]++
				}
			}
		}
	}
	for i, r := range ranks {
		for j, u := range r {
			if u == 0 {
				level = append(level, [2]int{i, j})
			}
		}
	}
	rank := 1
	for len(level) > 0 {
		var next [][2]int
		for _, ij := range level {
			i, j := ij[0], ij[1]
			ranks[i][j] = rank
			for _, xy := range [...][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
				x, y := xy[0], xy[1]
				if x >= 0 && y >= 0 && x < n && y < m && matrix[x][y] > matrix[i][j] {
					ranks[x][y]--
					if ranks[x][y] == 0 {
						next = append(next, [2]int{x, y})
					}
				}
			}
		}
		level = next
		rank++
	}
	return ranks
}
