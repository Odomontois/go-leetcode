package diceRollSimulation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiceSimulator(t *testing.T) {
	assert.Equal(t, 34, dieSimulator(2, []int{1, 1, 2, 2, 2, 3}))
	assert.Equal(t, 30, dieSimulator(2, []int{1, 1, 1, 1, 1, 1}))
	assert.Equal(t, 181, dieSimulator(3, []int{1, 1, 1, 2, 2, 3}))
}
