package main

import "fmt"

func main() {
	nums := []int{100, 2, 200, 4, 1, 3}
	fmt.Println(longestConsecutiveSequence(nums))
}

//[100,2,200,4,1,3]
func longestConsecutiveSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool, len(nums))
	// set all nums -> true
	for _, num := range nums {
		set[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !set[num-1] {
			curr := num
			length := 1
			for set[curr+1] {
				curr++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}
	return longest

}