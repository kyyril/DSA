package main

import (
	"fmt"
	"sort"
)


func main (){
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}
//find three sum = 0
func threeSum(nums []int)[][]int{
	sort.Ints(nums)
	res := [][]int{}

	for i:=0; i < len(nums); i++ {
		//skip duplicate i
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		right := len(nums)-1
		left := i+1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0{
				res = append(res, []int {nums[i], nums[left], nums[right]})
				//skip duplicate
				left++
				right--
				for left < right && nums[left] == nums[left-1]{
					left++
				}
				for left < right && nums[right] == nums[right+1]{
					right--
				}
			}else if sum < 0 {
				left++
			}else{
				right--
			}
		}
	}
	return res
}
