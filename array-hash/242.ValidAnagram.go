package main
import "fmt"

func main (){
	s, t :=  "anagram", "nagarap"
	result := isAnagram(s,t)
	fmt.Println(result)
}

func isAnagram (s string, t string) bool{
	if len(s) != len(t) {
		return false
	}
	freq := make(map[rune] int)

	for _, v := range s {
		freq[v]++
	}

	for _, v := range t {
		freq[v]--

		if freq[v] < 0 {
			return false
		}
	}

	return true
}