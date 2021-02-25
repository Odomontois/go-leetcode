package hard

var MOD Mnum = 1_000_000_007

type Mnum int64

func (x Mnum) Plus(y Mnum) Mnum {
	return (x + y) % MOD
}

func (x Mnum) Mul(y Mnum) Mnum {
	return (x * y) % MOD
}

func (x Mnum) Divide(y Mnum) Mnum {
	yr, _ := Euclid(y, MOD)
	return (x * (yr + MOD)) % MOD
}

func Euclid(a Mnum, b Mnum) (x Mnum, y Mnum) {
	if b == 0 {
		if a == 1 {
			return 1, 0
		}
		return 0, 0
	}
	z, x := Euclid(b, a%b)
	y = z - (a/b)*x
	return
}
