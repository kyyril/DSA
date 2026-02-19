package main

import "fmt"


func main (){
	nums := []int {1, 2, 3, 4, 5, 6, 7, 8}
	last := len(nums)-1
	fmt.Println(nums[:last])
} 