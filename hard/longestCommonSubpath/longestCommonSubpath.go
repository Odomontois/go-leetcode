package main

const P1 uint64 = 6_580_998_354_230_635_109
const P2 uint64 = 9_103_937_613_107_960_501

func hashes(path []int, k int) []uint64 {
	var h, hb uint64 = 0, P2
	res := []uint64{}
	for i, x := range path {
		h = h*P1 + uint64(x)*P2
		if i+1 < k {
			hb *= P1
			continue
		}
		res = append(res, h)
		h -= hb * uint64(path[i-k])
	}
	return res
}

type u struct{}

type HHSet map[uint64]u

func add(hhs HHSet, batch []uint64) HHSet {
	res := make(HHSet)
	for _, x := range batch {
		if hhs != nil {
			if _, ok := res[x]; !ok {
				continue
			}
		}
		res[x] = u{}
	}
	return res
}

func longestCommonSubpath(n int, paths [][]int) int {

}
