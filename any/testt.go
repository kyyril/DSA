package main
import "fmt"


func main() {
	numCourses := 2
	prerequisites := [][]int{{0,1}}
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		course := pre[0]
		prereq := pre[1]
		graph[prereq] = append(graph[prereq], course)
	}
	fmt.Println(graph)
}