package main

import "fmt"

type Default[A any] interface {
	*A
	Default() A
}

func GetDefault[A any, B Default[A]]() A {
	var b B
	return b.Default()
}

type Rational struct{ numerator, denominator int }

func (*Rational) Default() Rational {
	return Rational{0, 1}
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.numerator, r.denominator)
}

func main() {
	fmt.Println(GetDefault[Rational]())
}
