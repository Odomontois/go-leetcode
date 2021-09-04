package sets

import (
	"math/rand"
	"testing"
)

var input []uint64
var result []uint64

func prepareInput(b *testing.B) {
	if input != nil {
		return
	}
	b.StopTimer()
	for i := 0; i < 1000000; i++ {
		input = append(input, uint64(rand.Intn(30000)))
	}
	b.StartTimer()
}

type tt struct{}

func BenchmarkBoolMap(b *testing.B) {
	prepareInput(b)
	for i := 0; i < b.N; i++ {
		itemSet := make(map[uint64]bool)

		items := input
		for _, i := range items {
			itemSet[i] = true
		}
		items = nil
		for i := range itemSet {
			items = append(items, i)
		}
		result = items
	}
}
func BenchmarkStructMap(b *testing.B) {
	prepareInput(b)
	for i := 0; i < b.N; i++ {
		itemSet := make(map[uint64]tt)

		items := input
		for _, i := range items {
			itemSet[i] = tt{}
		}
		items = nil
		for i := range itemSet {
			items = append(items, i)
		}
		result = items
	}
}

func BenchmarkIFMap(b *testing.B) {
	prepareInput(b)
	for i := 0; i < b.N; i++ {
		itemSet := make(map[uint64]interface{})

		items := input
		for _, i := range items {
			itemSet[i] = true
		}
		items = nil
		for i := range itemSet {
			items = append(items, i)
		}
		result = items
	}
}
