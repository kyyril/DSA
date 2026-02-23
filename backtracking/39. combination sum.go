package main

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	var dfs func(start int, total int)
	dfs = func(start int, total int) {
		// BASE CASE 1: Berhasil mencapai target
		if total == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// BASE CASE 2: Lewat dari target (Optimization)
		if total > target {
			return
		}

		// LOOP: Mencoba setiap kandidat angka
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])

			// KUNCI: Pakai 'i', bukan 'i + 1', supaya angka yang sama bisa dipake lagi
			dfs(i, total+candidates[i])

			// BACKTRACK: Hapus angka terakhir untuk mencoba angka berikutnya di loop
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return result
}