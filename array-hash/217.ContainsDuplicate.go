package main

import "fmt"


func main() {
	nums := []int{2, 7, 8, 9, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{}, len(nums))

    for _, n := range nums {
        if _,ok := seen[n]; ok {
			fmt.Println(seen)
            return true
        }
        seen[n] = struct{}{}
    }
    return false
}