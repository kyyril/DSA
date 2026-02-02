package main

import "fmt"
func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSumII(numbers,26))
}

// use two pointer becouse arr is sorted
func twoSumII(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}