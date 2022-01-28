package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	tt := kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8)
	assert.Equal(t, 13, tt)
}
