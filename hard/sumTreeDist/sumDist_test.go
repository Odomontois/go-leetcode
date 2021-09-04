package sumtreedist_test

import (
	sumtreedist "leetcode/hard/sumTreeDist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTreeDist(t *testing.T) {
	assert.Equal(t,
		[]int{8, 12, 6, 10, 10, 10},
		sumtreedist.SumOfDistancesInTree(
			6,
			[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
		))
}
