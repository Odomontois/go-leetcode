package hard

import "testing"

func Test_goodTriplets(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{{
		"test1",
		args{
			[]int{2, 0, 1, 3},
			[]int{0, 1, 2, 3},
		},
		1,
	}, {
		"test 2",
		args{
			[]int{4, 0, 1, 3, 2},
			[]int{4, 1, 0, 2, 3},
		},
		4,
	},
		{
			"test 4",
			args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			120,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := goodTriplets(tt.args.nums1, tt.args.nums2); gotRes != tt.wantRes {
				t.Errorf("goodTriplets() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
