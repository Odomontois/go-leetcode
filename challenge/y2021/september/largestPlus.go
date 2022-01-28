package september

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minall(a int, xs ...int) int {
	for _, x := range xs {
		a = min(a, x)
	}
	return a
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = minall(i+1, j+1, n-i, n-j)
		}
	}
	good := func(x int) bool {
		return x >= 0 && x < n
	}
	draw := func(axis, step int, mine []int) {
		var p [2]int
		copy(p[:], mine)
		for cur := 0; good(p[0]) && good(p[1]); cur++ {
			z := &matrix[p[0]][p[1]]
			if *z > cur {
				*z = cur
			}
			p[axis] += step
		}
	}

	for _, mine := range mines {
		matrix[mine[0]][mine[1]] = 0
		draw(0, 1, mine)
		draw(0, -1, mine)
		draw(1, 1, mine)
		draw(1, -1, mine)
	}
	best := 0
	for _, m := range matrix {
		for _, c := range m {
			if c > best {
				best = c
			}
		}
	}
	return best
}
