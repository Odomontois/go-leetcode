package hard

import (
	"testing"
)

func Test_minimumDifference(t *testing.T) {
	type args []int
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: []int{3, 9, 7, 3},
			want: 2,
		},
		{
			name: "Example2",
			args: []int{-36, 36},
			want: 72,
		},
		{
			name: "Example3",
			args: []int{2, -1, 0, 4, -2, -9},
			want: 0,
		},
		{
			name: "WA1",
			args: []int{25, 49, 39, 42, 57, 35},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
