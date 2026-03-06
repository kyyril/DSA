package main

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{}
	for _, w := range words {
		curr := root
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if curr.children[idx] == nil {
				curr.children[idx] = &TrieNode{}
			}
			curr = curr.children[idx]
		}
		curr.word = w
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)
	res := []string{}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			dfs(r, c, board, root, &res)
		}
	}
	return res
}

func dfs(r, c int, board [][]byte, node *TrieNode, res *[]string) {
	//boundary check && and nil check
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return
	}
	char := board[r][c]
	if char == '#' || node.children[char-'a'] == nil {
		return
	}
	//move to node children
	node = node.children[char-'a']
	if node.word != "" { // store if word is found
		*res = append(*res, node.word)
		node.word = "" //remove so that there no duplicates
	}
	board[r][c] = '#'
	// 4-way exploration
	dfs(r+1, c, board, node, res) //down
	dfs(r-1, c, board, node, res) //up
	dfs(r, c+1, board, node, res) //right
	dfs(r, c-1, board, node, res) //left

	//BACKTRACK: Return the letters to their original state
	board[r][c] = char
}