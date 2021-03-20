package randomPointCircle

import (
	"math"
	"math/rand"
	"sort"
)

type Solution struct {
	radius, xc, yc float64
}

func Constructor(radius, xc, yc float64) Solution {
	return Solution{radius, xc, yc}
}

func (sol *Solution) RandPoint() []float64 {
	phi := rand.Float64() * 2 * math.Pi
	r := math.Sqrt(rand.Float64()) * sol.radius
	return []float64{sol.xc + r*math.Cos(phi), sol.yc + r*math.Sin(phi)}
}

type Person struct{ created int }

type PersonByCreated Person

func Lol() {
	persons := make([]Person, 0)
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].created < persons[j].created
	})
}
