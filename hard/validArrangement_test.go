package hard

import (
	"reflect"
	"sort"
	"testing"
)

func Test_validArrangement(t *testing.T) {

	tests := []struct {
		name string
		args [][]int
	}{
		{
			"test1",
			[][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}},
		},
		{
			"test2",
			[][]int{{1, 3}, {3, 2}, {2, 1}},
		},
		{
			"test with loop",
			[][]int{{1, 2}, {2, 4}, {4, 2}, {2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validArrangement(tt.args)
			check(t, tt.args, got)
		})
	}
}

func sortSlice2(k [][]int) {
	sort.Slice(k, func(i, j int) bool {
		return k[i][0] < k[j][0] || k[i][0] == k[j][0] && k[i][1] < k[j][1]
	})
}

func check(t *testing.T, args, got [][]int) {
	for i := range got {
		if i > 0 {
			if got[i-1][1] != got[i][0] {
				t.Errorf("unlinked at %v %v", i, got)
			}
		}
	}
	sortSlice2(args)
	sortSlice2(got)
	if !reflect.DeepEqual(args, got) {
		t.Errorf("got different elements %v %v", args, got)
	}
}
