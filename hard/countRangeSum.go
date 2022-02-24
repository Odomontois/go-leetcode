package hard

import (
	"fmt"
	"sort"
)

func countRangeSum(nums []int, lower int, upper int) int {
	fen := makeFen(len(nums))
	s := 0
	for i, num := range nums {
		nums[i] = s
		s += num
	}
	nums = append(nums, s)
	sums := make([]int, len(nums))
	copy(sums, nums)
	sort.Ints(sums)
	res := 0
	for _, sum := range nums {
		from := sort.SearchInts(sums, sum-upper)
		until := sort.SearchInts(sums, sum-lower+1)
		res += fen.calc(from, until)
		cur := sort.SearchInts(sums, sum)
		fen.add(cur, 1)
	}
	return res
}

// Fen is a fenwick tree
type Fen struct {
	size  int
	elems []int
}

func makeFen(size int) Fen {
	return Fen{size, make([]int, size*3)}
}

func (fen Fen) add(index int, amt int) {
	var walk func(int, int, int)
	walk = func(start int, end int, p int) {
		fen.elems[p] += amt
		if end-start == 1 {
			return
		}
		m := (end + start + 1) / 2
		if index < m {
			walk(start, m, 2*p+1)
		} else {
			walk(m, end, 2*p+2)
		}
	}
	walk(0, fen.size, 0)
}

func (fen Fen) calc(from int, until int) int {
	var walk func(int, int, int) int
	walk = func(start int, end int, p int) int {
		if start >= until || end <= from {
			return 0
		}
		if start >= from && end <= until {
			return fen.elems[p]
		}
		m := (end + start + 1) / 2
		return walk(start, m, 2*p+1) + walk(m, end, 2*p+2)
	}
	return walk(0, fen.size, 0)
}

func main() {
	fmt.Println(5 / 2)
	// fmt.Println(countRangeSum([]int{1, 1, 1, 1, 1, 1}, 1, 1))
	// fmt.Println(countRangeSum([]int{1, 1, 1, 1, 1, 1}, 2, 2))
	// fmt.Println(countRangeSum([]int{1, 1, 1, 1, 1, 1}, 3, 3))
	// fmt.Println(countRangeSum([]int{1, 1, 1, 1, 1, 1}, 1, 3))
	// fmt.Println(countRangeSum([]int{1, -1, 1, -1, 1, -1}, 1, 1))
	// fmt.Println(countRangeSum([]int{1, -1, 1, -1, 1, -1}, -1, -1))
	// fmt.Println(countRangeSum([]int{1, -1, 1, -1, 1, -1}, 0, 0))
	// fmt.Println(countRangeSum([]int{1, -1, 1, -1, 1, -1}, -1, 1))
	// fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
