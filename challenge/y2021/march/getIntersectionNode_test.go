package march_test

import (
	"leetcode/challenge/y2021/march"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Node = march.Node

func TestGetIntesrsectionNode(t *testing.T) {
	common := Node(8, Node(4, Node(5, nil)))
	nodeA := Node(4, Node(1, common))
	nodeB := Node(5, Node(6, Node(1, common)))
	res := march.GetIntersectionNode(nodeA, nodeB)
	assert.Same(t, res, common)
}
