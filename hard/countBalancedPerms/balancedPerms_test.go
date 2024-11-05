package countbalancedperms

import (
	"fmt"
	"testing"
)

func Test_countBalancedPermutationsOld(t *testing.T) {
	tests := []struct {
		name string
		num  string
		want int
	}{
		{
			name: "Example 1",
			num:  "123",
			want: 2,
		},
		{
			name: "Example 2",
			num:  "112",
			want: 1,
		},
		{
			name: "Example 3",
			num:  "12345",
			want: 0,
		},
		{
			name: "test 1",
			num:  "1212121212121212121213",
			want: 2134440,
		},
		{
			name: "TLE 1",
			num:  "14394152469836346236005574117",
			want: 0,
		},
		{
			name: "TLE 2",
			num:  "22431051238906437161377495483547",
			want: 351680872,
		},
		{
			name: "TLE 3",
			num:  "936743669957054147457096237781383667457517817480",
			want: 113880417,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalancedPermutationsOld(tt.num); got != tt.want {
				t.Errorf("countBalancedPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reci(t *testing.T) {
	tests := []int64{
		1, 5, 7, 1324, 40000000,
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintln(tt), func(t *testing.T) {
			if got := (reci(tt) * tt) % MOD; got != 1 {
				t.Errorf("reci(%v) * %v = %v, not 1", tt, tt, got)
			}
		})
	}
}

func Test_countBalancedPermutations(t *testing.T) {
	tests := []struct {
		name string
		num  string
		want int
	}{
		{
			name: "Example 1",
			num:  "123",
			want: 2,
		},
		{
			name: "Example 2",
			num:  "112",
			want: 1,
		},
		{
			name: "Example 3",
			num:  "12345",
			want: 0,
		},
		{
			name: "test 1",
			num:  "1212121212121212121213",
			want: 2134440,
		},
		{
			name: "TLE 1",
			num:  "14394152469836346236005574117",
			want: 0,
		},
		{
			name: "TLE 2",
			num:  "22431051238906437161377495483547",
			want: 351680872,
		},
		{
			name: "TLE 3",
			num:  "9103224235015855909098892653211286258658354677013618352",
			want: 996692047,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalancedPermutationsCringe(tt.num); got != tt.want {
				t.Errorf("countBalancedPermutations(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
