package march

import (
	"leetcode/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TNode = data.TNode

func TestFlipBinaryTree(t *testing.T) {
	assert.Equal(t, []int{-1}, flipMatchVoyage(TNode(1, nil, TNode(2, nil, nil)), []int{2, 1}))
	assert.Equal(t, []int{1}, flipMatchVoyage(TNode(1, TNode(2, nil, nil), TNode(3, nil, nil)), []int{1, 3, 2}))
	assert.Equal(t, []int{}, flipMatchVoyage(TNode(1, TNode(2, nil, nil), TNode(3, nil, nil)), []int{1, 2, 3}))

}
