package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushDominoes(t *testing.T) {
	assert.Equal(t, "RR.L", pushDominoes("RR.L"))
	assert.Equal(t, "LL.RR.LLRRLL..", pushDominoes(".L.R...LR..L.."))
}
