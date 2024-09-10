package y2024

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		"example 1",
		args{commands: []int{4, -1, 3}, obstacles: [][]int{}},
		25,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}

}
