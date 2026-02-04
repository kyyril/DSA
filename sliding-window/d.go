package main

import "fmt"

func main() {
	prices := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(prices))
}

func lengthOfLongestSubstring(s string) int {
    left := 0
    maxLen := 0
    seen := make(map[byte] bool)
    for right:=0; right < len(s); right++ {
        charRight := s[right]

        for seen[charRight]{
            charLeft := s[left]
            delete(seen, charLeft)
            left++
        }

        seen[charRight] = true
        currLen := right - left+1
        if currLen > maxLen {
            maxLen = currLen
        }
    }
    return maxLen
}

// Right pointer always moves forward
// Left pointer moves only when duplicate exists