package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLength(t *testing.T) {
	assert.Equal(t, 3, findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
