package main

import "fmt"

func main() {
	var xs [10]struct{ int }
	var y struct{ int }

	campari(func(int) *struct{ int } { return new(struct{ int }) })
	campari(func(int) *struct{ int } { return &struct{ int }{0} })
	campari(func(i int) *struct{ int } { return &xs[i] })
	campari(func(int) *struct{ int } { return &y })

}

func campari(f func(int) *struct{ int }) {
	fmt.Println(f(1) == f(2))
}
