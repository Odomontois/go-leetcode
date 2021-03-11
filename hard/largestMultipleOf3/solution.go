package largestMultipleOf3

import (
	"sort"
	"strings"
)

func largestMultipleOfThree(digits []int) string {
	sort.Ints(digits)
	sum := 0
	digs := [...][2]int{{-1, -1}, {-1, -1}, {-1, -1}}
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] % 3
		sum += d
		digs[d] = [...]int{digits[i], digs[d][0]}
	}
	sum %= 3
	rem := [...]int{-1, -1}
	switch {
	case digits[len(digits)-1] == 0:
		return "0"
	case sum == 0:
	case digs[sum][0] > 0:
		rem[0] = digs[sum][0]
	case digs[3-sum][1] > 1:
		rem = digs[3-sum]
	default:
		return ""
	}
	var builder strings.Builder
	for i := len(digits) - 1; i >= 0; i-- {
		switch digits[i] {
		case rem[0]:
			rem[0] = -1
		case rem[1]:
			rem[1] = -1
		default:
			builder.WriteRune('0' + rune(digits[i]))
		}
	}
	return builder.String()
}

func LargestMultipleOfThree(digits []int) string {
	return largestMultipleOfThree(digits)
}
