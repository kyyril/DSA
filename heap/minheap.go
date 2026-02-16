package main

type MinHeap []int

// [10,20,30,40,50,60]
func (h MinHeap) parent(i int) int { return (i - 1) / 2 }
func (h MinHeap) left(i int) int   { return 2*i + 1 }
func (h MinHeap) right(i int) int  { return 2*i + 2 }

// heapify up -> check if less than root -> swap
func (h *MinHeap) heapifyUp(i int) {
	for (*h)[i] < (*h)[h.parent(i)] {
		(*h)[i], (*h)[h.parent(i)] = (*h)[h.parent(i)], (*h)[i]
	}
}

func main() {

}