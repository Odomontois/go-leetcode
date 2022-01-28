package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFincClosesElements(t *testing.T) {
	assert.Equal(t, []int{3, 3, 4}, findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
	assert.Equal(t, []int{4, 5, 6}, findClosestElements([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 5))
}
