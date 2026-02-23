package main

func subsets(nums []int) [][]int {
	var result [][]int
	var path []int

	var dfs func(start int)
	dfs = func(start int) {
		//step1: save each of valid path
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		//step2: try another number
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}