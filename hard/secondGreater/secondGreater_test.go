package secondgreater

import (
	"reflect"
	"testing"
)

func Test_secondGreaterElement(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{{
		"example 1",
		[]int{2, 4, 0, 9, 6},
		[]int{9, 6, 6, -1, -1},
	}, {
		"example 2",
		[]int{3, 3},
		[]int{-1, -1},
	}, {
		"wa 1",
		[]int{1, 17, 18, 0, 18, 10, 20, 0},
		[]int{18, 18, -1, 10, -1, -1, -1, -1},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondGreaterElement(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secondGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
