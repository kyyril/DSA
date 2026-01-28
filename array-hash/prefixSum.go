package main

import "fmt"
func main() {
	nums := []int{2, 2, 4, 4}
	fmt.Println(prefixSum(nums))
}

func prefixSum(nums []int) []int {
	prefix := make([]int, len(nums))

	prefix[0] = nums[0] //store early idx(2)
	for i:=1; i < len(nums); i++ {
		prefix[i] = prefix[i - 1] + nums[i]
		//prev prefix + current nums
	}
	return prefix
}

// When to use?

// Keywords:
// subarray sum
// range sum
// continuous part