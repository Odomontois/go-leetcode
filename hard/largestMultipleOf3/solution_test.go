package largestMultipleOf3_test

import (
	"leetcode/hard/largestMultipleOf3"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple(t *testing.T) {
	assert.Equal(t, "981", largestMultipleOf3.LargestMultipleOfThree([]int{8, 1, 9}))
	assert.Equal(t, "8760", largestMultipleOf3.LargestMultipleOfThree([]int{8, 6, 7, 1, 0}))
	assert.Equal(t, "", largestMultipleOf3.LargestMultipleOfThree([]int{1}))
	assert.Equal(t, "0", largestMultipleOf3.LargestMultipleOfThree([]int{0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, "", largestMultipleOf3.LargestMultipleOfThree([]int{5, 8}))
}
