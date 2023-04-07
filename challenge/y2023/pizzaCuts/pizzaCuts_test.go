package pizzacuts

import "testing"

func Test_ways(t *testing.T) {
	tests := []struct {
		name  string
		pizza []string
		k     int
		want  int
	}{{
		name:  "Example 1",
		pizza: []string{"A..", "AAA", "..."},
		k:     3,
		want:  3,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ways(tt.pizza, tt.k); got != tt.want {
				t.Errorf("ways() = %v, want %v", got, tt.want)
			}
		})
	}
}
