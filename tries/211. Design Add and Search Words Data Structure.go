type WordDictionary struct {
    children [26]*WordDictionary
    isEnd    bool
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
    curr := this
    for _, char := range word {
        index := char - 'a'
        if curr.children[index] == nil {
            curr.children[index] = &WordDictionary{}
        }
        curr = curr.children[index]
    }
    curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    // Kita panggil fungsi helper rekursif mulai dari index 0
    return this.dfs(0, word)
}

func (this *WordDictionary) dfs(index int, word string) bool {
    curr := this
    
    for i := index; i < len(word); i++ {
        char := word[i]
        
        if char == '.' {
            // BACKTRACKING: Coba semua 26 kemungkinan huruf
            for _, child := range curr.children {
                if child != nil {
                    // Jika jalur ini (mulai dari huruf selanjutnya) mengembalikan true
                    if child.dfs(i+1, word) {
                        return true
                    }
                }
            }
            // Jika setelah mencoba semua 26 huruf tidak ada yang cocok
            return false
        } else {
            // Kasus huruf biasa (a-z)
            idx := char - 'a'
            if curr.children[idx] == nil {
                return false
            }
            curr = curr.children[idx]
        }
    }
    
    return curr.isEnd
}