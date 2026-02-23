package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	path := []int{}

	var dfs func(int)

	dfs = func(start int) {
		fmt.Println(path)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
}
