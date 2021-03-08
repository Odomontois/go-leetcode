package chunkedPalindrome_test

import (
	"leetcode/hard/chunkedPalindrome"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestChunkedPalindrome(t *testing.T) {
	var pal = chunkedPalindrome.LongestDecomposition

	assert.Equal(t, 5, pal("ababa"))
	assert.Equal(t, 7, pal("ghiabcdefhelloadamhelloabcdefghi"))
	assert.Equal(t, 1, pal("merchant"))
	assert.Equal(t, 11, pal("antaprezatepzapreanta"))
}
