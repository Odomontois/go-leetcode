package main

func findIntegers(n int) int {
	p := 1
	for p < n {
		p *= 2
	}
	return calc(n, p)
}

var cache = make(map[int]int)

func calc(n, p int) int {
	if n == 0 {
		return 1
	} else if n < p {
		return calc(n, p/2)
	} else if cached, ok := cache[n]; ok {
		return cached
	}
	u := n & ^p
	if u >= p/2 && u != 0 {
		u = p/2 - 1
	}
	res := calc(p-1, p/2) + calc(u, p/4)
	cache[n] = res
	return res
}
