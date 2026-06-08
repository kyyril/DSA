package main

import "fmt"
// type Node struct {
// 	Value int
// 	Prev *Node
// 	Next *Node
// }

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

// Head
// [10]
// Tail

func (dll *DoubleList) Prepend(value int) {
	newNode := &Node{Value : value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	}else {
		newNode.Next = dll.head //inisiasi di depan newNode jadi head lama
		dll.head.Prev = newNode //lalu prev head lama diisi oleh newNode
		dll.head = newNode //now replace head is newNode
	}
	dll.size++
}
func (dll *DoubleList) Append(value int) {
	newNode := &Node{Value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	}else {
		newNode.Prev = dll.tail
		dll.tail.Next = newNode
		dll.tail = newNode
	}
	dll.size++
}
func (dll *DoubleList) Delete(value int) bool{
	cur := dll.head

	for cur != nil {
		if cur.Value == value {
			cur.Prev.Next = cur.Next
			cur.Next.Prev = cur.Prev
			dll.size--
			return true
		}
		cur = cur.Next
	}
	return false
}
func (dll *DoubleList) DisplayForward() {
	if dll.head == nil {
		fmt.Println("List kosong")
		return
	}

	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// func main() {
// 	dll := &DoubleList{}

// 	fmt.Println("--- Mengisi Double Linked List ---")
// 	dll.Append(10)
// 	dll.Append(20)
// 	dll.Append(30)
// 	dll.Prepend(5)

// 	fmt.Print("Cetak Maju  : ")
// 	dll.DisplayForward() // Output: 5 <-> 10 <-> 20 <-> 30 <-> nil
// 	fmt.Printf("Ukuran list : %d\n", dll.size)
	
// 	dll.Delete(20)
// 	dll.DisplayForward() // Output: 5 <-> 10 <-> 30 <-> nil
// 	fmt.Printf("Ukuran list : %d\n", dll.size)
// }