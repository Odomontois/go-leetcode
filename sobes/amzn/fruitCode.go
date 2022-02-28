package amzn

func fruitCode(codes [][]string, cart []string) bool {
	enc := makeEncoder()
	cartC := enc.encodeSlice(cart)
	matched := 0

	for _, code := range codes {
		var index kmp
		codeC := enc.encodeSlice(code)
		index.build(codeC)
		i := index.search(codeC, cartC[matched:])
		if i == -1 {
			return false
		}
		matched += i
	}
	return true
}

type encoder struct {
	codeMap map[string]int
	next    int
}

const (
	anything = -1
)

func makeEncoder() encoder {
	return encoder{make(map[string]int), 1}
}

func (enc *encoder) encodeSlice(slice []string) []int {
	var res []int
	for _, s := range slice {
		code := enc.next
		if s == "anything" {
			code = anything
		} else if c, ok := enc.codeMap[s]; ok {
			code = c
		} else {
			enc.next++
			enc.codeMap[s] = code
		}
		res = append(res, code)
	}
	return res
}

type kmp []int

func (index *kmp) calc(pattern, target []int, build bool) int {
	if len(pattern) == 0 {
		return 0
	}
	p := 0
	if build {
		*index = append(*index, 0)
		target = target[1:]
	}
	for i, x := range target {
		for p != 0 && x != pattern[p] && pattern[p] != anything {
			p = (*index)[p-1]
		}
		if x == pattern[p] || pattern[p] == anything {
			p++
		}
		if build {
			*index = append(*index, p)
		} else if p == len(pattern) {
			return i + 1
		}
	}
	return -1
}

func (index *kmp) build(pattern []int) {
	index.calc(pattern, pattern, true)
}

func (index *kmp) search(pattern, target []int) int {
	return index.calc(pattern, target, false)
}
