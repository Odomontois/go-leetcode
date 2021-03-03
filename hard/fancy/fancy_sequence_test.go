package fancy_test

import (
	"leetcode/hard/fancy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFancy(t *testing.T) {
	fancy := fancy.Constructor()
	fancy.Append(2)                        // fancy sequence: [2]
	fancy.AddAll(3)                        // fancy sequence: [2+3] -> [5]
	fancy.Append(7)                        // fancy sequence: [5, 7]
	fancy.MultAll(2)                       // fancy sequence: [5*2, 7*2] -> [10, 14]
	assert.Equal(t, 10, fancy.GetIndex(0)) // return 10
	fancy.AddAll(3)                        // fancy sequence: [10+3, 14+3] -> [13, 17]
	fancy.Append(10)                       // fancy sequence: [13, 17, 10]
	fancy.MultAll(2)                       // fancy sequence: [13*2, 17*2, 10*2] -> [26, 34, 20]
	assert.Equal(t, 26, fancy.GetIndex(0)) // return 26
	assert.Equal(t, 34, fancy.GetIndex(1)) // return 34
	assert.Equal(t, 20, fancy.GetIndex(2)) // return 20
	assert.Equal(t, -1, fancy.GetIndex(3)) // return 20
}
