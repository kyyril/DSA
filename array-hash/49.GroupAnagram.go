package main

import "fmt"
func main (){
	strs := []string {"eat","tea","tan","ate","nat","bat"}

	fmt.Println(groupAnagrams(strs))
}
// ["eat",'tea"]
// "eat" = a1,b1,c1
// same pattern same group
// use map to store each string
// key is pattern then value is each word that pattern
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int] []string, len(strs))

	for _, s := range strs {
		var count [26]int // a-z
		for i:= 0; i < len(s); i++ {
			count[s[i]-'a']++
		}
		groups[count] = append(groups[count], s)
		// store -> [[1,0,1..] : "eat"]
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}