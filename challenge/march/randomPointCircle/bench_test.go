package randomPointCircle_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type Solution1 struct {
	radius, left, top float64
}

func Constructor1(radius, xc, yc float64) Solution1 {
	return Solution1{radius, xc - radius, yc - radius}
}

func (sol *Solution1) RandPoint() (float64, float64) {
	r := sol.radius
	r2 := r * r

	for {
		x := rand.Float64() * 2 * r
		y := rand.Float64() * 2 * r
		if (x-r)*(x-r)+(y-r)*(y-r) <= r2 {
			return x + sol.left, y + sol.top
		}
	}
}

type Solution2 struct {
	radius, xc, yc float64
}

func Constructor2(radius, xc, yc float64) Solution2 {
	return Solution2{radius, xc, yc}
}

const Pi2 = math.Pi * 2

func (sol *Solution2) RandPoint() (float64, float64) {
	phi := rand.Float64() * Pi2

	r := math.Sqrt(rand.Float64()) * sol.radius
	sin := math.Cos(phi)
	cos := math.Sqrt(1 - sin*sin)
	if phi > math.Pi {
		cos = -cos
	}
	return sol.xc + r*cos, sol.yc + r*sin
}

type Solution3 struct {
	radius, xc, yc float64
}

func Constructor3(radius, xc, yc float64) Solution3 {
	return Solution3{radius, xc, yc}
}

func (sol *Solution3) RandPoint() (float64, float64) {
	r := rand.Float64() + rand.Float64()
	if r > 1 {
		r = 2 - r
	}
	phi := rand.Float64() * Pi2

	r *= sol.radius
	sin := math.Cos(phi)
	cos := math.Sqrt(1 - sin*sin)
	if phi > math.Pi {
		cos = -cos
	}
	return sol.xc + r*cos, sol.yc + r*sin
}

func BenchmarkCircle(b *testing.B) {

	count := 1000000

	b.Run(fmt.Sprintf("Rejection sampling, %d points", count), func(b *testing.B) { // 0.0690 ns/op

		sol := Constructor1(543, 123.2, -1232)

		for i := 0; i < count; i++ {
			sol.RandPoint()
		}

	})

	b.Run(fmt.Sprintf("Polar coordinates, %d points", count), func(b *testing.B) { // 0.0910 ns/op
		sol := Constructor2(543, 123.2, -1232)

		for i := 0; i < count; i++ {
			sol.RandPoint()
		}
	})

	b.Run(fmt.Sprintf("Polar Convolution, %d points", count), func(b *testing.B) { // 0.0910 ns/op
		sol := Constructor3(543, 123.2, -1232)

		for i := 0; i < count; i++ {
			sol.RandPoint()
		}
	})

}
