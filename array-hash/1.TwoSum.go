package main

import "fmt"


func main() {
	nums := []int{2, 7, 8, 9, 10}
	target := 10
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, n := range nums {
        need := target - n
        if idx, ok := seen[need]; ok {
            return []int{idx, i}
        }
        seen[n] = i
    }
    return []int{}
}
