package main

import (
	"leetcode/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeBst1(t *testing.T) {
	assert.Nil(t, canMerge([]*TreeNode{
		data.TNode(5, data.TLeaf(3), data.TLeaf(8)),
		data.TNode(3, data.TLeaf(2), data.TLeaf(6)),
	}))
}
