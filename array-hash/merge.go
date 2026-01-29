package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 8, 3}
	arr2 := []int{4, 5, 6}
	fmt.Println(mergeSortedArrays(arr1, arr2))
}

// Merge two sorted arrays
func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	// Sort both arrays first
	sort.Ints(arr1)
	sort.Ints(arr2)

	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	// Append remaining elements
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}
