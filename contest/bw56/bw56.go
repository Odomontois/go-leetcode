package main

import "math"

func countTriples(n int) int {
	res := 0
	for i := 3; i <= n; i++ {
		for j := i + 1; i*i+j*j <= n*n; j++ {
			s := i*i + j*j
			x := int(math.Floor(math.Sqrt(float64(s))))
			if x*x == s {
				res += 2
			}
		}
	}
	return res
}

func nearestExit(maze [][]byte, entrance []int) int {
	q := [][3]int{{entrance[0], entrance[1], 0}}
	was := make([][]bool, len(maze))
	for i := range was {
		was[i] = make([]bool, len(maze[0]))
	}
	was[entrance[0]][entrance[1]] = true
	for len(q) > 0 {
		next := q[0]
		q = q[1:]
		for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			x, y := next[0]+d[0], next[1]+d[1]
			if x < 0 || x >= len(maze) || y < 0 || y >= len(maze[0]) {
				if next[2] > 0 {
					return next[2]
				}
			} else if maze[x][y] == byte('.') && !was[x][y] {
				q = append(q, [3]int{x, y, next[2] + 1})
				was[x][y] = true
			}
		}
	}
	return -1
}

func sumGame(num string) bool {
	sum := 0
	counts := 0
	for i, c := range num {
		if c == '?' {
			if i*2 >= len(num) {
				counts--
			} else {
				counts++
			}
			continue
		}
		x := int(c - '0')
		if i*2 >= len(num) {
			sum -= x
		} else {
			sum += x
		}
	}

	return (counts%2 != 0) || sum != -9*(counts/2)
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	ways := make([][][2]int, len(passingFees))
	cache := make(map[[2]int]int)
	for _, e := range edges {
		ways[e[1]] = append(ways[e[1]], [2]int{e[0], e[2]})
		ways[e[0]] = append(ways[e[0]], [2]int{e[1], e[2]})
	}
	var find func(int, int) int
	find = func(to, time int) int {
		if to == 0 {
			return passingFees[0]
		}
		res, ok := cache[[2]int{to, time}]
		if ok {
			return res
		}
		res = -1
		for _, w := range ways[to] {
			if w[1] > time {
				continue
			}

			path := find(w[0], time-w[1])
			if path == -1 {
				continue
			}
			path += passingFees[to]

			if res == -1 || res > path {
				res = path
			}
		}

		cache[[2]int{to, time}] = res
		return res
	}

	return find(len(passingFees)-1, maxTime)
}
