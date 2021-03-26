package pizza3n

func maxSizeSlices(slices []int) int {
	pizza := Pizza3n{slices, map[[2]int]int{}, map[[2]int]int{}}
	best := 0
	for i := range slices {
		max(&best, pizza.Full(i, len(slices)))
	}
	return best
}

type Pizza3n struct {
	slices        []int
	except1, full map[[2]int]int
}

func max(acc *int, x int) {
	if *acc < x {
		*acc = x
	}
}

func (pizza *Pizza3n) Full(from, size int) int {
	if size == 0 || size%3 != 0 {
		return 0
	}
	res, ok := pizza.full[[2]int{from, size}]
	if !ok {
		for l := 1; l < size; l += 3 {
			max(&res, pizza.slices[pizza.ix(from+l)]+pizza.Except1(from, l)+pizza.Except1(pizza.ix(from+l+1), size-l-1))
		}
		pizza.full[[2]int{from, size}] = res
	}
	return res
}

func (pizza *Pizza3n) Except1(from, size int) int {
	if size == 1 || size%3 != 1 {
		return 0
	}
	res, ok := pizza.except1[[2]int{from, size}]
	if !ok {
		for l := 0; l < size; l += 3 {
			max(&res, pizza.Full(from, l)+pizza.Full(pizza.ix(from+l+1), size-l-1))
		}
		pizza.except1[[2]int{from, size}] = res
	}
	return res
}

func (pizza *Pizza3n) ix(i int) int {
	return i % len(pizza.slices)
}
