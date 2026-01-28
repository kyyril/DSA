package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 3}
	fmt.Println(hasDuplicate(nums))
}

func hasDuplicate(nums []int) bool {
	has := make(map[int]bool)

	for _, n := range nums {
		if has[n] {
			return true
		}
		has[n] = true
	}
	return false
}

// Pattern 2: Seen Before (Duplicate Check)
// When to use?

// duplicate
// unique
// contains
// seen before