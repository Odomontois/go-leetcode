package shortestSuperstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestSuperstring(t *testing.T) {
	checkSolution(t, 16, []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"})
}

func checkSolution(t *testing.T, expLen int, words []string) {
	result := shortestSuperstring(words)
	for _, w := range words {
		assert.Contains(t, result, w)
	}
	assert.Equal(t, expLen, len(result), "expected length differs")
}
