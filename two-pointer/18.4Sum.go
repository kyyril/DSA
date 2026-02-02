package main

import (
	"fmt"
	"sort"
)

func main (){
	nums := []int{2,2,2,2,2}
	fmt.Println(fourSum(nums, 8))
}

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    n := len(nums)
    if n < 4 { return res }
    for i := 0; i < n-3; i++ {
        if i > 0 && nums[i] == nums[i-1] { continue }
        // --- PRUNING i ---
        if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target { break }
        if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target { continue }
        for j := i + 1; j < n-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] { continue }
            // --- PRUNING j ---
            if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target { break }
            if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target { continue }
            left, right := j+1, n-1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum == target {
                    res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    right--
                    for left < right && nums[left] == nums[left-1] { left++ }
                    for left < right && nums[right] == nums[right+1] { right-- }
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return res
}