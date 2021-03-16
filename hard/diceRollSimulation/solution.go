package diceRollSimulation

const MOD = 1_000_000_007

func dieSimulator(n int, rollMax []int) int {
	init := func() [6][]int {
		xs := [6][]int{}
		for dice := range xs {
			xs[dice] = make([]int, rollMax[dice])
		}
		return xs
	}
	add := func(res *int, amt int) {
		*res = (*res + amt) % MOD
	}
	prev, next := init(), init()
	for dice := range prev {
		prev[dice][0] = 1
	}
	for step := 1; step < n; step++ {
		for dice, p := range prev {
			for i, k := range p {
				if i+1 < rollMax[dice] {
					add(&next[dice][i+1], k)
				}
				for d2 := 0; d2 < 6; d2++ {
					if d2 != dice {
						add(&next[d2][0], k)
					}
				}
			}
		}
		next, prev = prev, next
		for _, d := range next {
			for k := range d {
				d[k] = 0
			}
		}
	}
	var res int

	for _, p := range prev {
		for _, k := range p {
			add(&res, k)
		}
	}
	return res
}
