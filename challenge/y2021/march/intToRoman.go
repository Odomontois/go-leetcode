package march

import (
	"sort"
	"strings"
)

var digMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var digs []int

func intToRoman(num int) string {
	if digs == nil {
		for k := range digMap {
			digs = append(digs, k)
		}
		sort.Ints(digs)
	}

	var builder strings.Builder

	for num > 0 {
		idx := sort.SearchInts(digs, num)
		if idx == len(digs) || num < digs[idx] {
			idx--
		}
		num -= digs[idx]
		builder.WriteString(digMap[digs[idx]])
	}
	return builder.String()
}
