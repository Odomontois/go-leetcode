package february

func scoreOfParentheses(S string) int {
	stack := []int{0}
	for _, c := range S {
		switch c {
		case '(':
			stack = append(stack, 0)
		case ')':
			n := len(stack) - 1
			x := stack[n] * 2
			stack = stack[0:n]
			if x == 0 {
				x = 1
			}
			stack[n-1] += x

		}
	}
	return stack[0]
}
