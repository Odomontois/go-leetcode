package painthouse

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Example 1",
		args: args{
			houses: []int{0, 0, 0, 0, 0},
			cost:   [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}},
			m:      5,
			n:      2,
			target: 3,
		},
		want: 9,
	}, {
		name: "Example 2",
		args: args{
			houses: []int{0, 2, 1, 2, 0},
			cost:   [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}},
			m:      5,
			n:      2,
			target: 3,
		},
		want: 11,
	},
		{
			name: "Example 3",
			args: args{
				houses: []int{3, 1, 2, 3},
				cost:   [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
				m:      4,
				n:      3,
				target: 3,
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.houses, tt.args.cost, tt.args.m, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
