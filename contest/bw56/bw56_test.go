package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumGame(t *testing.T) {
	assert.Condition(t, func() bool { return !sumGame("?3295???") })
}
