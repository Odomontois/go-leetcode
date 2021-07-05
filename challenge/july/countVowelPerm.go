package main

const Mod int = 1000_000_007

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for k := 1; k < n; k++ {
		a, e, i, o, u = (e+i+u)%Mod, (a+i)%Mod, (e+o)%Mod, i, (i+o)%Mod
	}
	return (a + e + i + o + u) % Mod
}
