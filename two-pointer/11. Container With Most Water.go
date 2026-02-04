package main

import (
	"fmt"
)


func main (){
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}

//[1,8,6,2,5,4,8,3,7]

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    max := 0

    for left < right {
        h := height[left]
        if height[right] < h {
            h = height[right]
        }

        area := h * (right - left)
        if area > max {
            max = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return max
}