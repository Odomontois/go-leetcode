package countanagrams

const MOD int64 = 1_000_000_007

func euclid(x, y int64) (int64, int64) {
	if y == 1 {
		return 0, 1
	}
	u, v := euclid(y, x%y)
	return v, u - (x/y)*v
}

func rec(x int64) int64 {
	_, u := euclid(MOD, x)
	return (u + MOD) % MOD
}

func countAnagrams(s string) int {
	res, count, counts := int64(1), 0, [26]int{}
	for _, c := range s {
		if c == ' ' {
			count, counts = 0, [26]int{}
			continue
		}
		x := int(c - 'a')
		counts[x]++
		count++
		res = (res * int64(count)) % MOD
		res = (res * rec(int64(counts[x]))) % MOD
	}
	return 0
}
