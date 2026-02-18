package main

import "fmt"

type MaxHeap struct {
	array []int
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }
func (h MaxHeap) swap(i1 int, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// insert add element to heap
func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}
// extract the largest key and remove it from heap
func (h *MaxHeap) Extract()int{
	extracted := h.array[0]
	last := len(h.array)-1
	//when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract!")
		return -1
	}
	//take last index and put the root
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)
	return extracted
}

//maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array)-1
	l, r := left(index), right(index)
	childToCompare := 0
	//loop while index has a least one child
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		}else if h.array[l] > h.array[r] {//when left child is larger
			childToCompare = l
		}else {//when right child is larger
			childToCompare = r
		}
		//compare array value of current idx to larger child and swap if smaller than child
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		}else{
			return
		}
	}
}

func main(){
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10,20,30}
	for _,v := range buildHeap {
		m.insert(v)
		fmt.Println(m)
	}

	for i:=0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}