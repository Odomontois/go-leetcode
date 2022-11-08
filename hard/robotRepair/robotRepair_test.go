package robotrepair

import "testing"

func Test_minimumTotalDistance(t *testing.T) {
	type args struct {
		robots    []int
		factories [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{{
		"example 2",
		args{
			robots:    []int{1, -1},
			factories: [][]int{{-2, 1}, {2, 1}},
		},
		2,
	}, {
		"example 1",
		args{
			robots:    []int{0, 4, 6},
			factories: [][]int{{2, 2}, {6, 2}},
		},
		4,
	}, {
		"test 1",
		args{
			robots:    []int{0, 4, 6, 1, 7, 3, 2, 16, 8, 9, 11},
			factories: [][]int{{2, 10}, {6, 10}, {3, 1}, {7, 1}},
		},
		25,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalDistance(tt.args.robots, tt.args.factories); got != tt.want {
				t.Errorf("minimumTotalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
