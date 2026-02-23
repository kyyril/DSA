package main

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var backtrack func(r, c, index int) bool
	backtrack = func(r, c, index int) bool {
		//find all word
		if index == len(word) {
			return true
		}
		//check boundary and match word
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[index] {
			return false
		}
		//mark the cell (so it is not used twice in one line)
		temp := board[r][c]
		board[r][c] = '#'
		// 4-way exploration
		found := backtrack(r+1, c, index+1) || //down
			backtrack(r-1, c, index+1) || //up
			backtrack(r, c+1, index+1) || //right
			backtrack(r, c-1, index+1) //left

		board[r][c] = temp
		return found
	}

	//starting point
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}

	return false
}