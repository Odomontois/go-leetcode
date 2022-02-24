package hard

import (
	"reflect"
	"testing"
)

func Test_findAllPeople(t *testing.T) {
	type args struct {
		n           int
		meetings    [][]int
		firstPerson int
	}

	a2, r2 := genTest2()
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	"test1",
		// 	args{6,
		// 		[][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}},
		// 		1},
		// 	[]int{0, 1, 2, 3, 5},
		// },
		{
			"big test",
			args{100000,
				a2,
				1},
			r2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllPeople(tt.args.n, tt.args.meetings, tt.args.firstPerson); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func genTest2() ([][]int, []int) {
	var arg [][]int
	res := []int{0}
	for i := 1; i < 100000; i++ {
		arg = append(arg, []int{0, i, 1})
		res = append(res, i)
	}
	return arg, res
}
