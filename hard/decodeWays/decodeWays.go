package main

func code(fst, snd int) int {
	i := fst*11 + snd
	if i > 40 {
		return 41
	}
	return i
}

func numDecodings(s string) int {
	const Mod int = 1_000_000_007
	var cc [42]int

	for c := 10; c <= 26; c++ {
		fst, snd := c/10+1, c%10+1
		cc[code(fst, snd)]++
		cc[code(0, snd)]++
		if snd > 1 {
			cc[code(fst, 0)]++
			cc[code(0, 0)]++
		}
	}
	cur, prev, cprev := 1, 0, 0

	for i, ch := range s {
		var next, c int
		switch ch {
		case '*':
			next = cur * 9
			c = 0
		case '0':
			next = 0
			c = 1
		default:
			next = cur
			c = int(ch-'0') + 1
		}

		if i > 0 && cprev != 1 {
			next += prev * cc[code(cprev, c)]
		}
		cur, prev, cprev = next%Mod, cur, c
	}
	return cur
}
