package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	enc := Constructor([]byte("abcd"), []string{"ei", "zf", "ei", "am"}, []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"})
	assert.Equal(t, enc.Encrypt("abcd"), "eizfeiam")
	assert.Equal(t, enc.Decrypt("eizfeiam"), 2)
}
