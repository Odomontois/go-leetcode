package y2024

import (
	"leetcode/data"
	"testing"
)

var TNode, TLeaf = data.TNode, data.TLeaf

func Test_countPairs(t *testing.T) {

	tests := []struct {
		name     string
		root     *TreeNode
		distance int
		want     int
	}{
		{
			name:     "example 1",
			root:     TNode(1, TNode(2, nil, TLeaf(4)), TLeaf(3)),
			distance: 3,
			want:     1,
		},
		{
			name:     "example 2",
			root:     TNode(1, TNode(2, TLeaf(4), TLeaf(5)), TNode(3, TLeaf(6), TLeaf(7))),
			distance: 3,
			want:     2,
		},
		{
			name:     "example 3",
			root:     TNode(7, TNode(1, TLeaf(6), nil), TNode(4, TLeaf(5), TNode(3, nil, TLeaf(2)))),
			distance: 3,
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.root, tt.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
