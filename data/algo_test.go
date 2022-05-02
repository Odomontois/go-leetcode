package data

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"},
		MapSlice([]int{1, 2, 3}, strconv.Itoa))
}
