// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/

package longestpath

import (
	"testing"
)

func Test_longestPath(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		"Example 1",
		args{
			parent: []int{-1, 0, 0, 1, 1, 2},
			s:      "abacbe",
		},
		3,
	}, {
		"Example 2",
		args{
			parent: []int{-1, 0, 0, 0},
			s:      "aabc",
		},
		3,
	}, {
		"WA 1",
		args{
			parent: []int{-1, 56, 10, 79, 52, 0, 37, 39, 127, 125, 116, 52, 95, 131, 105, 55, 55, 52, 87, 35, 43, 130, 87, 103, 8, 73, 8, 116, 4, 43, 60, 104, 116, 118, 78, 9, 133, 139, 7, 127, 96, 28, 52, 79, 78, 36, 102, 134, 100, 104, 47, 127, 129, 77, 121, 133, 10, 58, 104, 55, 69, 73, 107, 9, 139, 79, 52, 72, 130, 78, 112, 43, 14, 4, 120, 9, 139, 118, 52, 52, 73, 82, 79, 58, 121, 80, 139, 10, 25, 74, 10, 123, 134, 112, 40, 80, 108, 128, 5, 52, 43, 31, 10, 42, 79, 139, 86, 58, 3, 118, 117, 21, 4, 79, 45, 26, 5, 122, 102, 13, 88, 139, 108, 118, 116, 10, 58, 32, 80, 125, 121, 105, 116, 104, 82, 131, 39, 10, 126, 125},
			s:      "vodpyvpjmogqvwnibqasbulkfbfugvtlpdtsmydrbrekavkhifoypepbcnzpmasbnlrfqgdhnmvhldsrogjsntummchcftzrnycichziopmphfqwqdsihoywdpqqkyvzrhbqorwrkmns",
		},
		17,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPath(tt.args.parent, tt.args.s); got != tt.want {
				t.Errorf("longestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
