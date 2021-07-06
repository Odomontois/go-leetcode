package streamOfChars

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestStreamOfChars(t *testing.T) {
	checker := Constructor([]string{"cd", "f", "kl"})
	tests := "abcDeFghijkL"
	for _, c := range tests {
		assert.Equal(
			t,
			unicode.IsUpper(c),
			checker.Query(byte(unicode.ToLower(c))))
	}
}

func TestStreamOfChars2(t *testing.T) {
	checker := Constructor([]string{"cde", "cd", "def"})
	tests := "abcDEFghijklmn"
	for _, c := range tests {
		assert.Equal(
			t,
			unicode.IsUpper(c),
			checker.Query(byte(unicode.ToLower(c))),
			"char is %v",
			string(c))
	}
}

func TestStreamOfChars3(t *testing.T) {
	checker := Constructor([]string{"cde", "cd", "def", "klmn", "lm", "mn"})
	tests := "abcDEFghijklMN"
	for _, c := range tests {
		assert.Equal(
			t,
			unicode.IsUpper(c),
			checker.Query(byte(unicode.ToLower(c))),
			"char is %v",
			string(c))
	}
}
