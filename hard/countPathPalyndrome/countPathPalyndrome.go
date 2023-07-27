package countpathpalyndrome

func countPalindromePaths(parent []int, s string) int64 {
	children := make([][]int, len(parent))

	for i, p := range parent {
		if p >= 0 {
			children[p] = append(children[p], i)
		}
	}
	pathsCount := make(map[uint32]int64)

	var descend func(int, uint32)
	descend = func(node int, mask uint32) {
		pathsCount[mask]++
		for _, c := range children[node] {
			childMask := mask ^ (1 << (s[c] - 'a'))
			descend(c, childMask)
		}
	}
	descend(0, 0)
	res := int64(0)
	for mask, count := range pathsCount {
		res += count * (count - 1)
		for i := 0; i < 26; i++ {
			res += count * pathsCount[mask^(1<<i)]
		}
	}
	return res / 2
}
