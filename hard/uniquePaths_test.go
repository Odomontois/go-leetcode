package hard_test

import (
	"leetcode/hard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsIII(t *testing.T) {
	assert.Equal(t, 2,
		hard.UniquePathsIII([][]int{
			{1, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 2, -1}}))
}
