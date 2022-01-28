package february

// RomanToInt Cool thing
func RomanToInt(s string) int {
	return romanToInt(s)
}

func romanToInt(s string) int {
	cur := 0
	prev := 0
	acc := 0

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			cur = 1
		case 'V':
			cur = 5
		case 'X':
			cur = 10
		case 'L':
			cur = 50
		case 'C':
			cur = 100
		case 'D':
			cur = 500
		case 'M':
			cur = 1000
		}
		if cur < prev {
			acc -= cur
		} else {
			acc += cur
		}
		prev = cur
	}
	return acc
}

// func main() {
// 	fmt.Println(romanToInt("III"))
// 	fmt.Println(romanToInt("IV"))
// 	fmt.Println(romanToInt("IX"))
// 	fmt.Println(romanToInt("LVIII"))
// 	fmt.Println(romanToInt("MCMXCIV"))
// }

func brokenCalc(X int, Y int) int {
	res := 0
	for Y > X {
		res += Y%2 + 1
		Y = (Y + 1) / 2
	}

	return res + X - Y
}

// func main() {
// 	fmt.Println(brokenCalc(2, 3))
// 	fmt.Println(brokenCalc(5, 8))
// 	fmt.Println()
// }
