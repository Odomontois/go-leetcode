package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSetSize(t *testing.T) {
	assert.Equal(t, 2, minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}
