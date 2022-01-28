package main

func customSortString(order string, str string) string {
	cs := make(map[rune]int)
	for _, c := range str {
		cs[c]++
	}
	var res []rune
	for _, c := range order {
		for i := 0; i < cs[c]; i++ {
			res = append(res, c)
		}
		delete(cs, c)
	}
	for c, t := range cs {
		for i := 0; i < t; i++ {
			res = append(res, c)
		}
	}
	return string(res)
}
