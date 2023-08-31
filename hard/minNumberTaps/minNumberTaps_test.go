package minnumbertaps

import "testing"

func Test_minTaps(t *testing.T) {
	type args struct {
		n      int
		ranges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{5, []int{3, 4, 1, 1, 0, 0}},
			1,
		},
		{
			"example2",
			args{3, []int{0, 0, 0, 0}},
			-1,
		},
		{
			"test1",
			args{5, []int{2, 2, 2, 2, 2, 2}},
			2,
		},
		{
			"test2",
			args{5, []int{1, 1, 1, 1, 1, 1}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTaps(tt.args.n, tt.args.ranges); got != tt.want {
				t.Errorf("minTaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
