package main

func threeEqualParts(arr []int) []int {
	ones := 0
	for _, x := range arr {
		ones += x
	}
	if ones%3 != 0 {
		return []int{-1, -1}
	}
	ones /= 3
	if ones == 0 {
		return []int{0, len(arr) - 1}
	}
	x1, y1 := collectOnes(arr, ones, 0)
	x2, y2 := collectOnes(arr, ones, y1)
	x3, y3 := collectOnes(arr, ones, y2)
	y1 += len(arr) - y3
	y2 += len(arr) - y3
	if x2 >= y1 && x3 >= y2 && sliceEq(arr[x1:y1], arr[x2:y2]) && sliceEq(arr[x1:y1], arr[x3:]) {
		return []int{y1 - 1, y2}
	}
	return []int{-1, -1}
}

func sliceEq(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, a := range x {
		if a != y[i] {
			return false
		}
	}
	return true
}

func collectOnes(arr []int, ones, start int) (int, int) {
	i := start
	for arr[i] == 0 {
		i++
	}
	j := i
	for ones > 0 {
		ones -= arr[j]
		j++
	}
	return i, j
}
