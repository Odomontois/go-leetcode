package hashmap

type node struct {
	key  int
	val  int
	next *node
}

type MyHashMap struct {
	table []*node
	size  int
}

func Constructor() (hm MyHashMap) { return }

func (hm *MyHashMap) resize(newSize int) {
	newTable := make([]*node, newSize)
	for _, head := range hm.table {
		for ; head != nil; head = head.next {
			cell := head.key % newSize
			newTable[cell] = &node{head.key, head.val, newTable[cell]}
		}
	}
	hm.table = newTable
}

func (hm *MyHashMap) Put(key int, val int) {
	if len(hm.table) < hm.size+1 {
		hm.resize(2*hm.size + 1)
	}
	if node := hm.lookup(key); node != nil {
		node.val = val
		return
	}
	cell := key % len(hm.table)
	hm.table[cell] = &node{key, val, hm.table[cell]}
	hm.size++
}

func (hm *MyHashMap) Remove(key int) {
	if hm.size == 0 {
		return
	}
	cell := key % len(hm.table)
	ptr := &hm.table[cell]
	for cur := hm.table[cell]; cur != nil; ptr, cur = &cur.next, cur.next {
		if key == cur.key {
			*ptr = cur.next
			hm.size--
		}
	}
	if 4*hm.size < len(hm.table) {
		hm.resize(2 * hm.size)
	}
}

func (hm *MyHashMap) Get(key int) int {
	if node := hm.lookup(key); node != nil {
		return node.val
	}
	return -1
}

func (hm *MyHashMap) lookup(key int) *node {
	if hm.size == 0 {
		return nil
	}
	cell := key % len(hm.table)
	for cur := hm.table[cell]; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur
		}
	}
	return nil
}
