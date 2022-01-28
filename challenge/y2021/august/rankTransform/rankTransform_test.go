package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankTransform1(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {2, 3}},
		matrixRankTransform([][]int{{1, 2}, {3, 4}}))
}
func TestRankTransform2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 1}, {1, 1}},
		matrixRankTransform([][]int{{7, 7}, {7, 7}}))
}
func TestRankTransform3(t *testing.T) {

	assert.Equal(t, [][]int{{4, 2, 3}, {1, 3, 4}, {5, 1, 6}, {1, 3, 4}},
		matrixRankTransform([][]int{{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19}}))
}
func TestRankTransform4(t *testing.T) {

	assert.Equal(t, [][]int{{5, 1, 4}, {1, 2, 3}, {6, 3, 1}},
		matrixRankTransform([][]int{{7, 3, 6}, {1, 4, 5}, {9, 8, 2}}))
}
