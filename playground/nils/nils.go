package main

import "fmt"

func main() {
	var nil1 *int
	var nil2 *string
	var n1, n2, n3 any = nil, nil1, nil2
	for _, x := range []any{n1, n2, n3} {
		for _, y := range []any{n1, n2, n3} {
			fmt.Print(x == y, " ")
		}
		fmt.Println()
	}
}
