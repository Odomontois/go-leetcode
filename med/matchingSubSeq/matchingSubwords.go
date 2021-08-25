package matchingsubseq

func numMatchingSubseq(s string, words []string) int {
	const A = int('a')
	res := 0
	var let [26][][2]int
	for j, w := range words {
		c := int(w[0]) - A
		let[c] = append(let[c], [2]int{j, 0})
	}
	for _, r := range s {
		c := int(r) - A
		move := let[c]
		let[c] = [][2]int{}
		for _, d := range move {
			w := words[d[0]]
			d[1]++
			if d[1] >= len(w) {
				res++
			} else {
				c := int(w[d[1]]) - A
				let[c] = append(let[c], d)
			}
		}
	}
	return res
}
