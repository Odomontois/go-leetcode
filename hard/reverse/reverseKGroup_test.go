package reverse_test

import (
	"leetcode/hard"
	"leetcode/hard/reverse"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	assert.Equal(
		t, []int{3, 2, 1, 6, 5, 4, 7, 8},
		reverse.ReverseKGroup(hard.MakeList([]int{1, 2, 3, 4, 5, 6, 7, 8}), 3).ToSlice(),
	)
}
