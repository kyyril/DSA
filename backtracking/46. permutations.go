package main

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(nums) == len(path) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return result

}