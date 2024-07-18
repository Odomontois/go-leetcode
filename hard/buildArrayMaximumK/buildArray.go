package buildarraymaximumk

type caching[V any] interface {
	comparable
	precheck() (V, bool)
}

func debug(xs ...any) {
	// fmt.Println(xs...)
}

func Cached[K caching[V], V any](name string, calc func(K) V) func(K) V {
	cache := make(map[K]V)
	return func(key K) V {
		if v, ok := key.precheck(); ok {
			debug(name, "precheck", key, v)
			return v
		}
		if v, ok := cache[key]; ok {
			return v
		}
		res := calc(key)
		cache[key] = res
		debug(name, "result", key, res)
		return res
	}
}

type pk struct{ m, p int }

func (pk pk) precheck() (res int, ready bool) {
	switch {
	case pk.p == 0:
		return 1, true
	case pk.m <= 1:
		return pk.m, true
	default:
		return
	}
}

type build struct{ n, m, k int }

func (args build) precheck() (res int, ready bool) {
	switch {
	case args.k == 0 && args.n == 0:
		return 1, true
	case args.k == 0 || args.n == 0 || args.m == 0 || args.k > args.n:
		return 0, true
	case args.m == 1 && args.k == 1:
		return 1, true
	default:
		return
	}
}

const MOD int64 = 1_000_000_007

func mplus(x, y int) int {
	return int((int64(x) + int64(y)) % MOD)
}
func mmul(x, y int) int {
	return int((int64(x) * int64(y)) % MOD)
}

func numOfArrays(n int, m int, k int) int {
	var prod func(pk) int
	prod = Cached("prod", func(pk pk) int {
		pk.p--
		return mmul(prod(pk), pk.m)
	})
	var fexact, fsum func(build) int
	fexact = Cached("fexact", func(args build) int {
		args.n--
		resPrev := mmul(fexact(args), args.m)
		args.m--
		args.k--
		resCur := fsum(args)
		return mplus(resPrev, resCur)
	})
	fsum = Cached("fsum", func(args build) int {
		cur := fexact(args)
		args.m--
		prev := fsum(args)
		return mplus(cur, prev)
	})
	return fsum(build{n, m, k})
}
