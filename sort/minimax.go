package main

import "fmt"
func main (){
	nums := []int{1,2,9,3,4,5}
	fmt.Println(mini(nums))
	fmt.Println(max(nums))
}

func mini (nums []int) int {
	min := nums[0]
	for i:=1; i < len(nums); i++{
		if nums[i] < min{
			min = nums[i]
		}
	}
	return min
}
func max (nums []int) int {
	max := nums[0]
	for i:=1; i < len(nums); i++{
		if nums[i] > max{
			max = nums[i]
		}
	}
	return max
}