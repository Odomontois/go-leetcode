package hard

func waysToFillArray(queries [][]int) []int {
	res := make([]int, 0)
	fill := MakeFill(10_000, 10_000)
	for _, q := range queries {
		n, k := q[0], q[1]
		res = append(res, fill.ArrLen(n, k))
	}
	return res
}
func WaysToFillArray(queries [][]int) []int {
	return waysToFillArray(queries)
}

type Primes []int

func MakePrimes(n int) Primes {
	primes := []int{2, 3, 5}
	ks := [2]int{1, 5}
	for i := 6; i <= n; i += 6 {
		for _, k := range ks {
			d := i + k
			for _, p := range primes {
				if p*p > d {
					primes = append(primes, d)
					break
				}
				if d%p == 0 {
					break
				}
			}
		}
	}
	return primes
}

type Fact []Mnum

func MakeFact(n int) Fact {
	res := []Mnum{1}
	var cur Mnum = 1
	for k := 1; k <= n; k++ {
		cur = (cur * Mnum(k)) % MOD
		res = append(res, cur)
	}
	return res
}

func (fact Fact) Comb(n int, k int) Mnum {
	return fact[n].Divide(fact[k]).Divide(fact[n-k])
}

func (fact Fact) CombRep(n int, k int) Mnum {
	return fact.Comb(n+k-1, n)
}

type Fill struct {
	c      Fact
	primes Primes
}

func MakeFill(l int, n int) Fill {
	return Fill{c: MakeFact(l + 32), primes: MakePrimes(n)}
}

func (fill Fill) ArrLen(n int, k int) int {
	var res int64 = 1
	for _, d := range fill.primes.Divs(k) {
		res = (res * int64(fill.c.CombRep(d.count, n))) % int64(MOD)
	}
	return int(res)
}

type Div struct {
	div   int
	count int
}

func (ps Primes) Divs(n int) []Div {
	res := make([]Div, 0)
	for _, p := range ps {
		if p*p > n {
			break
		}
		count := 0
		for n%p == 0 {
			count++
			n /= p
		}
		if count > 0 {
			res = append(res, Div{p, count})
		}
	}
	if n > 1 {
		res = append(res, Div{n, 1})
	}
	return res
}
