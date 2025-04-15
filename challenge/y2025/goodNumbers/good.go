package goodintegers

func countGoodIntegers(n int, ks int) (res int64) {
	if n == 1 {
		return 9 / int64(ks)
	}

	fact := MakeFact()

	p := int64(1)
	for i := 0; i < n/2; i++ {
		p *= 10
	}
	k := int64(ks)
	md := int64(1 + (n%2)*9)
	vars := make(map[[10]int]bool)
	for x := p / 10; x < p; x++ {
		pxb := x*p*md + rev(x)
		for d := int64(0); d < md; d++ {
			px := pxb + d*p
			if px%k != 0 {
				continue
			}
			var ds [10]int
			for y := px; y > 0; y /= 10 {
				ds[y%10]++
			}
			vars[ds] = true
		}
	}
	for ds := range vars {
		res += fact.Count(n, ds)
	}
	return res
}

type Fact [11]int64

func MakeFact() Fact {
	var f Fact
	f[0] = 1
	for i := int64(1); i < 11; i++ {
		f[i] = f[i-1] * i
	}
	return f
}

func (f *Fact) Comb(n, k int) int64 {
	if n < k {
		return 0
	}
	return f[n] / f[k] / f[n-k]
}

func (f *Fact) Count(n int, ds [10]int) int64 {
	nd, count := n-ds[0], f.Comb(n-1, ds[0])
	for _, jd := range ds[1:] {
		count *= f.Comb(nd, jd)
		nd -= jd
	}
	return count
}

func rev(x int64) int64 {
	r := int64(0)
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}
