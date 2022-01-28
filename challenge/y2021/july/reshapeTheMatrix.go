package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat[0])*len(mat) {
		return mat
	}
	mid := make([]int, 0, r*c)
	res := make([][]int, 0, r)
	for _, row := range mat {
		mid = append(mid, row...)
	}
	for len(mid) > 0 {
		res = append(res, mid[0:c])
		mid = mid[c:]
	}
	return res
}
