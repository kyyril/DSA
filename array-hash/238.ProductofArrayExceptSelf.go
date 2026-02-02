package main

import "fmt"

func main() {
	nums := []int{-1,1,0,-3,3}
	fmt.Println(productExceptSelf(nums))
}

// [1,2,3,4]
// left  [1,1,2,6]
func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    // Left pass
    res[0] = 1
    for i := 1; i < n; i++ {
        res[i] = res[i-1] * nums[i-1]
    }

    // Right pass
    right := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]
    }

    return res
}
