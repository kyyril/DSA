package main

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &WordDictionary{}
		}
		curr = this.children[index]
	}
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	// Kita panggil fungsi helper rekursif mulai dari index 0
	return this.dfs(0, word)
}

func (this *WordDictionary) dfs(index int, word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if char == '.' {
			//backtracking: try all 26 char
			for _, child := range curr.children {
				
			}
		} else {
			index := char - 'a'
			if curr.children[index] == nil {
				return false
			}
			curr = curr.children[index]
		}
	}
	return curr.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */