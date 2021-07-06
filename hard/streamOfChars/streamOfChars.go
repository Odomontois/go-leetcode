package streamOfChars

type Node struct {
	chars    map[byte]*Node
	fallback *Node
	word     bool
}

func (node *Node) descend(c byte) *Node {
	next, ok := node.chars[c]
	for !ok && node.fallback != nil {
		node = node.fallback
		next, ok = node.chars[c]
	}
	if ok {
		return next
	}
	return node
}

func (node *Node) putChar(c byte, last bool) *Node {
	if node.chars == nil {
		node.chars = make(map[byte]*Node)
	} else if next, ok := node.chars[c]; ok {
		next.word = last || next.word
		return next
	}
	next := new(Node)
	if node.fallback != nil {
		next.fallback = node.fallback.descend(c)
	} else {
		next.fallback = node
	}
	next.word = last || next.fallback.word
	node.chars[c] = next
	return next
}

type StreamChecker struct {
	cur *Node
}

type wnode struct {
	word string
	node *Node
}

func Constructor(words []string) StreamChecker {
	root := new(Node)
	wnodes := make([]wnode, 0, len(words))
	for _, w := range words {
		wnodes = append(wnodes, wnode{w, root})
	}
	for len(wnodes) > 0 {
		wnodes2 := make([]wnode, 0, len(wnodes))
		for _, wn := range wnodes {
			last := len(wn.word) == 1
			next := wn.node.putChar(wn.word[0], last)
			if !last {
				wnodes2 = append(wnodes2, wnode{wn.word[1:], next})
			}
		}
		wnodes = wnodes2
	}
	return StreamChecker{root}
}

func (checker *StreamChecker) Query(letter byte) bool {
	checker.cur = checker.cur.descend(letter)
	return checker.cur.word
}
