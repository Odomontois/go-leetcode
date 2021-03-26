package pizza3n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPizza(t *testing.T) {
	assert.Equal(t, 10, maxSizeSlices([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 16, maxSizeSlices([]int{8, 9, 8, 6, 1, 1}))
	assert.Equal(t, 21, maxSizeSlices([]int{4, 1, 2, 5, 8, 3, 1, 9, 7}))
}
