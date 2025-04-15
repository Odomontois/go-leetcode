package goodintegers

import "testing"

func Test_countGoodIntegers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{
			name:    "Example 1",
			args:    args{n: 3, k: 5},
			wantRes: 27,
		},
		{
			name:    "Example 2",
			args:    args{n: 1, k: 4},
			wantRes: 2,
		},
		{
			name:    "Example 3",
			args:    args{n: 5, k: 6},
			wantRes: 2468,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countGoodIntegers(tt.args.n, tt.args.k); gotRes != tt.wantRes {
				t.Errorf("countGoodIntegers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
