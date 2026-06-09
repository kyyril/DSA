package main

import "fmt"

type Stack [T any] struct {
	elements []T
}
func(s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func(s *Stack[T]) Pop() (T, bool){
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	//take element
	lastIdx := len(s.elements)-1
	element := s.elements[lastIdx]

	//remove element
	s.elements = s.elements[:lastIdx]
	return element, true
}

func(s *Stack[T]) Peek() (T, bool){
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return s.elements[len(s.elements)-1], true
}

func(s *Stack[T]) IsEmpty() bool{
	return len(s.elements) == 0 
}

func(s *Stack[T]) Size() int{
	return len(s.elements)
}

func main() {
	stack := &Stack[int]{}
	stack.Push(10)
	stack.Push(15)
	stack.Push(20)
	fmt.Printf("size: %d\n",stack.Size())
	if top, ok := stack.Peek(); ok {
		fmt.Printf("\nTop: %d\n", top)
	}
	for !stack.IsEmpty(){
		item, _ := stack.Pop()
		fmt.Printf("Pop: %d\n", item)
	} 
	fmt.Printf("\nIs Empty? %t\n", stack.IsEmpty())
}