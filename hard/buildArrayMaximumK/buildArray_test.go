package buildarraymaximumk

import "testing"

func Test_numOfArrays(t *testing.T) {
	type args struct {
		n int
		m int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example 1",
			args{n: 2, m: 3, k: 1},
			6,
		},
		{
			"test 1",
			args{n: 2, m: 2, k: 2},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfArrays(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("numOfArrays(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
