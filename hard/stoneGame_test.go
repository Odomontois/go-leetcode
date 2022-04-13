package hard_test

import (
	"leetcode/hard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGame(t *testing.T) {
	t.Skip("Skip Stone Game")
	assert.Equal(t, 18, hard.StoneGameV([]int{6, 2, 3, 4, 5, 5}))
	assert.Equal(t, 330, hard.StoneGameV([]int{98, 77, 24, 49, 6, 12, 2, 44, 51, 96}))
}
