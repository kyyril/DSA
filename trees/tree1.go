package main

import "fmt"

type TreeNodes struct {
	Val   int
	Left  *TreeNodes
	Right *TreeNodes
}

//create node
//     1
//   /   \
//  2     3

func main() {
	// Create root node (1)
	root := &TreeNodes{Val: 1}

	// Create left child (2)
	root.Left = &TreeNodes{Val: 2}

	// Create right child (3)
	root.Right = &TreeNodes{Val: 3}
	
	fmt.Printf("Root: %d\n", root.Val)
	fmt.Printf("Left child: %d\n", root.Left.Val)
	fmt.Printf("Right child: %d\n", root.Right.Val)
}
