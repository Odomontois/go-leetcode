package trionic2

import "testing"

func Test_maxSumTrionic(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int64
	}{
		{
			name: "Example 1",
			nums: []int{0, -2, -1, -3, 0, 2, -1},
			want: -4,
		},
		{
			name: "Example 2",
			nums: []int{1, 4, 2, 7},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSumTrionic(tt.nums)
			if got != tt.want {
				t.Errorf("maxSumTrionic() = %v, want %v", got, tt.want)
			}
		})
	}
}
