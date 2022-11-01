package similararrays

import "testing"

func Test_makeSimilar(t *testing.T) {
	type args struct {
		nums   []int
		target []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{{
		"example 1",
		args{[]int{8, 12, 6}, []int{2, 14, 10}},
		2,
	}, {
		"example 2",
		args{[]int{1, 2, 5}, []int{4, 1, 3}},
		1,
	}, {
		"example 3",
		args{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		0,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := makeSimilar(tt.args.nums, tt.args.target); gotRes != tt.wantRes {
				t.Errorf("makeSimilar() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
