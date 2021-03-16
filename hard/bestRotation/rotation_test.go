package bestRotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestRotation(t *testing.T) {
	assert.Equal(t, 3, bestRotation([]int{2, 3, 1, 4, 0}))
	assert.Equal(t, 0, bestRotation([]int{1, 3, 0, 2, 4}))
}
