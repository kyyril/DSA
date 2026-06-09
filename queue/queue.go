package main

import "fmt"
type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// take and delete element from front of elements
// return element and condition
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	first := q.elements[0]
	// handle garbage collection memory with "zero out"
	// GC can clear memory if "T" is pointer to object (couse GC is automaticly clear object that are no longer use)
	var zeroValue T //without generic we can just use "nil" value
	q.elements[0] = zeroValue

	q.elements = q.elements[1:] //next
	return first, true
}

func (q *Queue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return q.elements[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	qiu := &Queue[string]{}
	qiu.Enqueue("wahyu")
	qiu.Enqueue("jamal")

	if front, ok := qiu.Front(); ok {
		fmt.Printf("\nFront: %s", front)
	}
	for !qiu.IsEmpty() {
		item,_ := qiu.Dequeue()
		fmt.Printf("\n dequeue: %s", item)
	}

	fmt.Printf("\nIs Empty: %t", qiu.IsEmpty())
}