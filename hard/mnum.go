package hard

// MOD ...
var MOD Mnum = 1_000_000_007

//Mnum is a modulo num
type Mnum int64

// Plus adds
func (x Mnum) Plus(y Mnum) Mnum {
	return (x + y) % MOD
}

// Mul multiplies
func (x Mnum) Mul(y Mnum) Mnum {
	return (x * y) % MOD
}

//Divide divides
func (x Mnum) Divide(y Mnum) Mnum {
	yr, _ := euclid(y, MOD)
	return (x * (yr + MOD)) % MOD
}

func euclid(a Mnum, b Mnum) (x Mnum, y Mnum) {
	if b == 0 {
		if a == 1 {
			return 1, 0
		}
		return 0, 0
	}
	z, x := euclid(b, a%b)
	y = z - (a/b)*x
	return
}
