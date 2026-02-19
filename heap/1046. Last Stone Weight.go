package main

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } //maxheap(big on top)
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// interface because we use library heap, name is type assertion
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int)) //.(int) force to int
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	*h = stones
	heap.Init(h)

	for h.Len() > 1 {
		stone1 := heap.Pop(h).(int) //big one
		stone2 := heap.Pop(h).(int) //second one
		if stone1 != stone2 {
			heap.Push(h, stone1 - stone2)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}
