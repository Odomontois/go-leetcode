package mindaysdisconnect

import "testing"

func Test_minDays(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{
			name: "water",
			args: [][]int{
				{0, 0},
				{0, 0}},
			want: 0,
		},
		{
			name: "one land",
			args: [][]int{
				{0, 1},
				{0, 0}},
			want: 1,
		},
		{
			name: "two islands",
			args: [][]int{
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 0}},
			want: 0,
		},
		{
			name: "de wurst",
			args: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{1, 1, 1}},
			want: 1,
		},
		{
			name: "H-shaped",
			args: [][]int{
				{1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1}},
			want: 1,
		},
		{
			name: "O-shaped",
			args: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1}},
			want: 2,
		},
		{
			name: "plain",
			args: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1}},
			want: 2,
		},
		{
			name: "example 1",
			args: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0}},
			want: 2,
		},
		{
			name: "example 2",
			args: [][]int{{1, 1}},
			want: 2,
		},
		{
			name: "A-shaped",
			args: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1}},
			want: 2,
		},
		{
			name: "dumbbell",
			args: [][]int{
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1}},
			want: 1,
		},
		{
			name: "two dots",
			args: [][]int{{1, 0, 1, 0}},
			want: 0,
		},
		{
			name: "peak",
			args: [][]int{
				{0, 1, 0, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
			},
			want: 1,
		},
		{
			name: "bridges",
			args: [][]int{
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1},
				{0, 1, 0, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "wa",
			args: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
