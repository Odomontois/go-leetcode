package hard

import "testing"

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		flowers    []int
		newFlowers int64
		target     int
		full       int
		partial    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			want: 14,
			args: args{
				flowers:    []int{1, 3, 1, 1},
				newFlowers: 7,
				target:     6,
				full:       12,
				partial:    1,
			},
		},
		{
			name: "Example 2",
			want: 30,
			args: args{
				flowers:    []int{2, 4, 5, 3},
				newFlowers: 10,
				target:     5,
				full:       2,
				partial:    6,
			},
		},
		{
			name: "WA 1",
			want: 41,
			args: args{
				flowers:    []int{8, 20, 13},
				newFlowers: 12,
				target:     12,
				full:       4,
				partial:    3,
			},
		},
		{
			name: "WA 2",
			want: 14,
			args: args{
				flowers:    []int{20, 1, 15, 17, 10, 2, 4, 16, 15, 11},
				newFlowers: 2,
				target:     20,
				full:       10,
				partial:    2,
			},
		},
		{
			name: "WA 3",
			want: 58,
			args: args{
				flowers:    []int{5, 5, 15, 1, 9},
				newFlowers: 36,
				target:     12,
				full:       9,
				partial:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.flowers, tt.args.newFlowers, tt.args.target, tt.args.full, tt.args.partial); got != tt.want {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
