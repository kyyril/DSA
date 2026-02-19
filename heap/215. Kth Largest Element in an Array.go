import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//nums:[3,2,1,5,6] k:2
//[1,2,3] -> pop[2,3]
//[2,3,5] -> pop[3,5]
//[3,5,6] -> pop[5,6]
//return k largest *h[0](5)
//time->0(nlogk) space->0(k)
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h) // Buang elemen terkecil jika size > k
		}
	}

	return (*h)[0] // Root min-heap adalah elemen terbesar ke-k
}