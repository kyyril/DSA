package main

import "fmt"


func main (){
	slice := []int{1, 2}
	slice = append(slice, 3, 4)
	newSlice := []int{5, 6}
	slice = append(slice, newSlice...)
	fmt.Println(slice) // [1 2 3 4 5 6]
	fmt.Println("====append=====")

	//copy
	goal := make([]int, len(slice))
	copy(goal, slice)
	fmt.Println(goal) // [1 2 3 4 5 6]
	fmt.Println("====copy=====")

	//len() & cap()
	s := make([]int, 3, 10)
	fmt.Println(len(s)) // 3
	fmt.Println(cap(s)) // 10
	fmt.Println("====len&cap=====")

	// remove i
	i := 2
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
	fmt.Println("====remove[i]=====")

}
