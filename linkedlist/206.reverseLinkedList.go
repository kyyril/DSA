package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	curr := head
	fmt.Print("[ ")
	for curr != nil {
		fmt.Printf("%d ", curr.Val) // Cetak angka saat ini
		curr = curr.Next            // Loncat ke node selanjutnya
	}
	fmt.Println("]")
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev 
		prev = curr
		curr = nextTemp
	}
	return prev // Mengembalikan head baru (ujung kanan yang sekarang jadi depan)
}

func main() {
	// Membuat Linked List: 0 -> 1 -> 2 -> 3 -> nil
	node3 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	head := &ListNode{Val: 0, Next: node1}

	// PRINT 1: Sebelum Dibalik
	fmt.Print("Sebelum di-reverse (Print jalan maju): ")
	printList(head)

	// JALANKAN REVERSE (Memutar semua isi variabel Next)
	newHead := reverseList(head)

	// PRINT 2: Setelah Dibalik
	fmt.Print("Setelah di-reverse  (Print jalan mundur): ")
	printList(newHead)
}