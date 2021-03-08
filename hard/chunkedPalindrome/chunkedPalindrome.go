package chunkedPalindrome

const HA uint64 = 490_832_341_198_073_269
const HB uint64 = 8_426_150_837_985_834_059

func longestDecomposition(text string) int {
	indices := make([][]uint64, 0)
	for i := range text {
		cur, suffix := uint64(0), make([]uint64, 0)
		for _, c := range text[i:] {
			cur = cur*HA + uint64(c) + HB
			suffix = append(suffix, cur)
		}
		indices = append(indices, suffix)
	}
	pals := make([]int, len(text)/2+1)
	pals[len(text)/2] = len(text) % 2
	for i := len(pals) - 2; i >= 0; i-- {
		best := 1
		for k := 0; k < len(pals)-i-1; k++ {
			mirror := len(text) - i - k - 1
			next := pals[i+k+1] + 2
			if indices[i][k] == indices[mirror][k] && next > best {
				best = next
			}
		}
		pals[i] = best
	}
	return pals[0]
}

type PrefNode struct {
	next     *int
	idx      int
	children []*PrefNode
}

func Prefix() *PrefNode {
	node := new(PrefNode)
	node.next = new(int)
	return node
}

func (node *PrefNode) push(char rune) *PrefNode {
	char -= 'a'
	if node.children == nil {
		node.children = make([]*PrefNode, 26)
	}
	down := node.children[char]
	if down == nil {
		down = new(PrefNode)
		(*node.next)++
		down.idx = *node.next
		down.next = node.next
		node.children[char] = down
	}
	return down
}

func LongestDecomposition(text string) int {
	return longestDecomposition(text)
}
