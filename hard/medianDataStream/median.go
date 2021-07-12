package main

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	smol MaxHeap
	big  MinHeap
}

type MinHeap struct {
	sort.IntSlice
}

func (mh MinHeap) Peek() int {
	return mh.IntSlice[0]
}

func (mh *MinHeap) Push(x interface{}) {
	mh.IntSlice = append(mh.IntSlice, x.(int))
}

func (mh *MinHeap) Pop() interface{} {
	sl := mh.IntSlice
	n := len(sl) - 1
	mh.IntSlice = sl[:n]
	return sl[n]
}

type MaxHeap struct {
	MinHeap
}

func (mh MaxHeap) Less(i int, j int) bool {
	return mh.IntSlice[i] > mh.IntSlice[j]
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}

}

func (finder *MedianFinder) AddNum(num int) {
	if finder.smol.Len() > 0 && finder.smol.Peek() < num {
		heap.Push(&finder.big, num)
		if finder.big.Len() > finder.smol.Len() {
			heap.Push(&finder.smol, heap.Pop(&finder.big))
		}
	} else {
		heap.Push(&finder.smol, num)
		if finder.smol.Len() > finder.big.Len()+1 {
			heap.Push(&finder.big, heap.Pop(&finder.smol))
		}
	}
}

func (finder *MedianFinder) FindMedian() float64 {
	if finder.smol.Len() > finder.big.Len() {
		return float64(finder.smol.Peek())
	}
	return float64(finder.smol.Peek()+finder.big.Peek()) / 2.0
}
