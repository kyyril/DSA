package main

import "fmt"

func main() {
	num := []int{1, 2, 2, 3, 3, 3}
	frequencyCount(num)
}

func frequencyCount (num []int){
	freq := make(map[int]int)
	for _, n := range num {
		freq[n]++
	}
	fmt.Println(freq)
}

// Pattern 1: Frequency Counting
// When to use?

// Keywords:
// count
// frequency
// how many times
// most common
// anagram