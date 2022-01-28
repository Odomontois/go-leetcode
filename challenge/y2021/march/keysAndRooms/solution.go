package keysAndRooms

func canVisitAllRooms(rooms [][]int) bool {
	seen, n := make([]bool, len(rooms)), len(rooms)
	seen[0] = true
	for i, stack := 0, []int{-1}; i >= 0; i, stack = stack[len(stack)-1], stack[:len(stack)-1] {
		n--
		for _, k := range rooms[i] {
			if !seen[k] {
				seen[k] = true
				stack = append(stack, k)
			}
		}
	}
	return n == 0
}
