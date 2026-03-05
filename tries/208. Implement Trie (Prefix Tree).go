package main

import "fmt"

type Node struct {
	children [26]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{}}
}

// create path one byone
func (this *Trie) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a' // Convert 'a'-'z' ke 0-25
		if curr.children[index] == nil {
			curr.children[index] = &Node{}
		}
		curr = curr.children[index]
	}
	curr.isEnd = true //mark as last char
}

// should be found the path and isEnd = true
func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a' // Convert 'a'-'z' ke 0-25
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	fmt.Println(curr.isEnd)
	return curr.isEnd
}

// just found the path, don't care about isEnd
//Selama rantai pointer menyambung, hasil pasti true.
func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a' // Convert 'a'-'z' ke 0-25
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	fmt.Println(true)
	return true
}

func main() {
	trie := &Trie{&Node{}}
	trie.Insert("apple")
	trie.Search("apple") // return True
	trie.Search("app")  // return False
	trie.StartsWith("app")// return True
	trie.Insert("app")
	trie.Search("app") // return True
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */