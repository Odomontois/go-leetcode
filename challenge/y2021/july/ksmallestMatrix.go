package main

import "container/heap"

type Elem struct {
	val int
	row int
}
type ElemHeap []Elem

func (eh ElemHeap) Len() int {
	return len(eh)
}

func (eh ElemHeap) Less(i int, j int) bool {
	return eh[i].val < eh[j].val
}

func (eh ElemHeap) Swap(i int, j int) {
	eh[i], eh[j] = eh[j], eh[i]
}

func (eh *ElemHeap) Push(x interface{}) {
	*eh = append(*eh, x.(Elem))
}

func (eh *ElemHeap) Pop() interface{} {
	n := len(*eh) - 1
	ret := (*eh)[n]
	*eh = (*eh)[:n]
	return ret
}

func (eh *ElemHeap) MatPush(m [][]int, row int) {
	heap.Push(eh, Elem{m[row][0], row})
	m[row] = m[row][1:]
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix[0])
	h := new(ElemHeap)
	h.MatPush(matrix, 0)
	for i := 0; i < k; i++ {
		el := heap.Pop(h).(Elem)
		if len(matrix[el.row]) > 0 {
			h.MatPush(matrix, el.row)
		}
		if el.row+1 < len(matrix) && len(matrix[el.row+1]) == n {
			h.MatPush(matrix, el.row+1)
		}
	}
	return heap.Pop(h).(Elem).val
}
