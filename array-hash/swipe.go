package main

import "fmt"

func main (){
	nums := []int {1,2,3,4,5}
	fmt.Println(nums)
	fmt.Println(swipeLeft(nums))
}

func swipeRight(nums []int) []int {
    n := len(nums)
    if n <= 1 {
        return nums
    }
    last := nums[n-1]

    // pointer from back
    for i := n - 1; i > 0; i-- {
        nums[i] = nums[i-1]
        fmt.Println(nums)
    }

    nums[0] = last

    return nums
}

// store first num
// [1,2,3] 
// [2,2,3] 0 
// [2,3,3] 1
// [2,3,1] 2

func swipeLeft (nums []int) []int {
    first := nums[0]
    for i := 0; i < len(nums)-1; i++ {
        nums[i] = nums[i+1] //copy value +i
    }
    nums[len(nums) -1 ] = first
    return nums
}