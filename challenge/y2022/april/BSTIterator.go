package april

import "leetcode/data"

type BSTIterator struct {
	ch   <-chan int
	next int
	read bool
}

func Constructor(root *data.TreeNode) BSTIterator {
	ch := make(chan int)
	go func() {
		treeIt(root, ch)
		close(ch)
	}()
	return BSTIterator{ch: ch}
}

func (it *BSTIterator) Next() int {
	it.HasNext()
	it.read = false
	return it.next
}

func (it *BSTIterator) HasNext() bool {
	if !it.read {
		it.next, it.read = <-it.ch
	}
	return it.read
}

func treeIt(root *data.TreeNode, vals chan<- int) {
	if root == nil {
		return
	}
	treeIt(root.Left, vals)
	vals <- root.Val
	treeIt(root.Right, vals)
}
