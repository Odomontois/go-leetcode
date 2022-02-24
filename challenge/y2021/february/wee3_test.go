package february_test

import (
	"leetcode/challenge/y2021/february"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var check = func(s string, exp int) {
		res := february.RomanToInt(s)
		if res != exp {
			t.Errorf("for %s got %v, want %v", s, res, exp)
		}
	}
	check("III", 3)
	check("IV", 4)
	check("IX", 9)
	check("LVIII", 58)
	check("MCMXCIV", 1994)
}
