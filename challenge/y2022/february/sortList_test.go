package february

import (
	"leetcode/data"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{{
		"test 1",
		[]int{4, 3, 2, 1},
		[]int{1, 2, 3, 4},
	}, {
		"test 2",
		[]int{-1, 5, 3, 4, 0},
		[]int{-1, 0, 3, 4, 5},
	}, {
		"test 3",
		nil,
		nil,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(data.List(tt.args)).ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
