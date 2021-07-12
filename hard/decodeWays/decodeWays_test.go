package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeWays(t *testing.T) {
	assert.Equal(t, 404, numDecodings("*1*1*0"))
	assert.Equal(t, 9, numDecodings("*"))
	assert.Equal(t, 18, numDecodings("1*"))
	assert.Equal(t, 15, numDecodings("2*"))
	assert.Equal(t, 999, numDecodings("***"))
	assert.Equal(t, 2203848, numDecodings("**2**3**4"))
}
