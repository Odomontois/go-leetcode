package main

func grayCode(n int) []int {
	res, cur := []int{0}, 0

	for i := 1; i < (1 << n); i++ {
		cur ^= (i ^ (i - 1) + 1) / 2
		res = append(res, cur)
	}
	return res
}
