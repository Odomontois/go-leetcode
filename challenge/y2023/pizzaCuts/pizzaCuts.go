package pizzacuts

const MOD int64 = 1_000_000_007

type point struct{ i, j int }

type key struct {
	from point
	k    int
}

func ways(pizza []string, k int) int {
	counts := make([][]int, len(pizza)+1)
	counts[0] = make([]int, len(pizza[0])+1)
	for i := range pizza {
		counts[i+1] = make([]int, len(pizza[i])+1)
		for j := range pizza[i] {
			cur := 0
			if pizza[i][j] == 'A' {
				cur = 1
			}
			counts[i+1][j+1] = cur + counts[i][j+1] + counts[i+1][j] - counts[i][j]
		}
	}

	has := func(from, to point) bool {
		return counts[to.i][to.j]-counts[from.i][to.j]-counts[to.i][from.j]+counts[from.i][from.j] > 0
	}

	cache := make(map[key]int)
	end := point{len(pizza), len(pizza[0])}

	var f func(from point, k int) int
	f = func(from point, k int) int {
		if !has(from, end) {
			return 0
		}
		if k == 1 {
			return 1
		}
		key := key{from, k}
		if v, ok := cache[key]; ok {
			return v
		}
		res := int64(0)
		for i := from.i + 1; i < end.i; i++ {
			if has(from, point{i, end.j}) {
				res += int64(f(point{i, from.j}, k-1))
			}
		}
		for j := from.j + 1; j < end.j; j++ {
			if has(from, point{end.i, j}) {
				res += int64(f(point{from.i, j}, k-1))
			}
		}

		resInt := int(res % MOD)
		cache[key] = resInt
		// fmt.Printf("cache[%v] = %v\n", key{from, k}, resInt)
		return resInt
	}

	return f(point{0, 0}, k)
}
