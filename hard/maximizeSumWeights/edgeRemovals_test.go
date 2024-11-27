package maximizesumweights

import "testing"

func Test_maximizeSumOfWeights(t *testing.T) {
	type args struct {
		edges [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example 1",
			args: args{edges: [][]int{{0, 1, 4}, {0, 2, 2}, {2, 3, 12}, {2, 4, 6}}, k: 2},
			want: 22,
		},
		{
			name: "example 2",
			args: args{edges: [][]int{{0, 1, 5}, {1, 2, 10}, {0, 3, 15}, {3, 4, 20}, {3, 5, 5}, {0, 6, 10}}, k: 3},
			want: 65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeSumOfWeights(tt.args.edges, tt.args.k); got != tt.want {
				t.Errorf("maximizeSumOfWeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
