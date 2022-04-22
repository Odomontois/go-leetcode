package hashset

type Node struct {
	val  int
	next *Node
}

type MyHashSet struct {
	table []*Node
	size  int
}

func Constructor() (set MyHashSet) { return }

func (hs *MyHashSet) resize(newSize int) {
	newTable := make([]*Node, newSize)
	for _, head := range hs.table {
		for ; head != nil; head = head.next {
			cell := head.val % newSize
			newTable[cell] = &Node{val: head.val, next: newTable[cell]}
		}
	}
	hs.table = newTable
}

func (hs *MyHashSet) Add(key int) {
	if len(hs.table) < hs.size+1 {
		hs.resize(2*hs.size + 1)
	}
	if hs.Contains(key) {
		return
	}
	cell := key % len(hs.table)
	hs.table[cell] = &Node{val: key, next: hs.table[cell]}
	hs.size++
}

func (hs *MyHashSet) Remove(key int) {
	if hs.size == 0 {
		return
	}
	cell := key % len(hs.table)
	ptr := &hs.table[cell]
	for cur := hs.table[cell]; cur != nil; ptr, cur = &cur.next, cur.next {
		if key == cur.val {
			*ptr = cur.next
			hs.size--
		}
	}
	if 4*hs.size < len(hs.table) {
		hs.resize(2 * hs.size)
	}
}

func (hs *MyHashSet) Contains(key int) bool {
	if hs.size == 0 {
		return false
	}
	cell := key % len(hs.table)
	for cur := hs.table[cell]; cur != nil; cur = cur.next {
		if cur.val == key {
			return true
		}
	}
	return false
}
