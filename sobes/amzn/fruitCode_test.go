package amzn

import (
	"strings"
	"testing"
)

func Test_fruitCode(t *testing.T) {
	type args struct {
		code [][]string
		cart []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{{
		"example 1",
		args{
			code: [][]string{{"apple", "apple"}, {"banana", "anything", "banana"}},
			cart: strings.Split("orange, apple, apple, banana, orange, banana", ", "),
		},
		true,
		false,
	}, {
		"example 2",
		args{
			code: [][]string{{"apple", "apple"}, {"banana", "anything", "banana"}},
			cart: strings.Split("banana, orange, banana, apple, apple", ", "),
		},
		false,
		false,
	}, {
		"example 3",
		args{
			code: [][]string{{"apple", "apple"}, {"banana", "anything", "banana"}},
			cart: strings.Split("apple, banana, apple, banana, orange, banana", ", "),
		},
		false,
		false,
	},
		{
			"example 4",
			args{
				code: [][]string{{"apple", "apple"}, {"apple", "apple", "banana"}},
				cart: strings.Split("apple, apple, apple, banana", ", "),
			},
			false,
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fruitCode(tt.args.code, tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("fruitCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fruitCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
