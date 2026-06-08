package main

import "fmt"


type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		//add to the first linkedlist element
		newNode.Next = l.head
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	for curr.Next != nil {
		curr = curr.Next //road to last element
	}
	curr.Next = newNode
	l.size++
}

func (l *LinkedList) Delete(value int) bool{
	if l.head == nil {
		return false
	}
	//if value is head just delete
	if value == l.head.Value {
		l.head = l.head.Next
		l.size--
		return true
	}
	//seacrh node with same value
	curr := l.head 
	for curr.Next != nil && curr.Next.Value != value {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next //replace with next node
		l.size--
		return true
	}
	return false
}

func (l *LinkedList) Display() {
	if l.head == nil {
		fmt.Println("LinkedList empty!")
		return
	}
	curr := l.head
	for curr != nil {
		fmt.Printf("%d->", curr.Value)
		curr = curr.Next
	}
	fmt.Println(nil)
}
func main (){
	list := &LinkedList{}
	list.Prepend(20)
	list.Prepend(10)
	list.Append(5)
	list.Display()
	list.Delete(5)
	list.Display()
}