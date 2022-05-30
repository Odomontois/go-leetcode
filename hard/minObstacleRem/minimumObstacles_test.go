package minobstaclerem

import "testing"

func Test_minimumObstacles(t *testing.T) {
	type args [][]int

	tests := []struct {
		name string
		args args
		want int
	}{{
		"example 1",
		[][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}},
		2,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumObstacles(tt.args); got != tt.want {
				t.Errorf("minimumObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
