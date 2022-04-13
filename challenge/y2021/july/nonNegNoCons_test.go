package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIntegers(t *testing.T) {
	t.Skip("Skip Find Integeres")
	check := func(exp, n int) {
		assert.Equal(t, exp, findIntegers(n), "n = %v", n)
	}
	check(0, 0)
	check(1, 1)
	check(2, 2)
	check(2, 3)
	check(3, 4)
	check(4, 5)
	check(4, 6)
	check(4, 7)
	check(5, 8)
	check(6, 9)
	check(7, 10)
	check(7, 11)
	check(7, 12)
	check(7, 13)
	check(7, 14)
	check(7, 15)
	check(8, 16)
}
