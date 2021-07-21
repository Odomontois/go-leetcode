package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func pushDominoes(dominoes string) string {
	var cur []int
	res := make([]rune, len(dominoes))
	for i, d := range dominoes {
		res[i] = d
		switch d {
		case 'L':
			cur = append(cur, -i-1)
		case 'R':
			cur = append(cur, i+1)
		default:
		}
	}

	for len(cur) > 0 {
		var next []int
		for _, i := range cur {
			n := i + 1
			if abs(n) == 0 || abs(n) > len(res) {
				continue
			}

			c := &res[abs(n)-1]
			switch *c {
			case '.':
				if n < 0 {
					*c = 'l'
				} else {
					*c = 'r'
				}
			case 'l':
				if n > 0 {
					*c = '.'
				}
			case 'r':
				if n < 0 {
					*c = '.'
				}
			default:
				continue
			}
			next = append(next, n)
		}
		cur = nil
		for _, i := range next {
			c := &res[abs(i)-1]
			switch *c {
			case 'l':
				*c = 'L'
			case 'r':
				*c = 'R'
			default:
				continue
			}
			cur = append(cur, i)
		}
	}

	return string(res)
}
