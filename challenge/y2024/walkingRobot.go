package y2024

func robotSim(commands []int, obstacles [][]int) (res int) {
	obsSet := make(map[[2]int]bool)
	for _, obs := range obstacles {
		obsSet[[2]int{obs[0], obs[1]}] = true
	}
	x, y, dx, dy := 0, 0, 0, 1
	for _, c := range commands {
		switch c {
		case -1:
			dx, dy = dy, -dx
		case -2:
			dx, dy = -dy, dx
		default:
			for ; c > 0; c-- {
				u, v := x+dx, y+dy
				if obsSet[[2]int{u, v}] {
					break
				}
				x, y = u, v
			}
			if d := x*x + y*y; d > res {
				res = d
			}
		}
	}
	return res
}
