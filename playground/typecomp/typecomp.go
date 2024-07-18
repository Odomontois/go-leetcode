package main

import "fmt"

type A interface {
	X(string) int
}
type B interface {
	X(string) int
}

type C struct {
	a A
}

type D struct {
	a B
}

type E = struct {
	a A
}

func check(a A, c C) []any {
	var b B = a
	var d D = D(c)
	var e E = c
	return []any{
		b,
		d,
		e,
	}
}

func main() {
	fmt.Println("Hello, 世界")
}
