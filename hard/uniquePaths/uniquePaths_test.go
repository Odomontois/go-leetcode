package uniquePaths_test

import (
	"leetcode/hard/uniquePaths"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsIII(t *testing.T) {
	assert.Equal(t, 2,
		uniquePaths.UniquePathsIII([][]int{
			{1, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 2, -1}}))
}
