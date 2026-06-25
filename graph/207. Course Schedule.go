package main
import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		course := pre[0]
		prereq := pre[1]
		graph[prereq] = append(graph[prereq], course)
	}

	state := make([]int, numCourses)
	var hasCycle func(course int) bool
	hasCycle = func(course int) bool{
		if state[course] ==  1 {
			return true
		}
		if state[course] == 2 {
			return false
		}

		state[course] = 1
		for _, nextCourse := range graph[course]{
			if hasCycle(nextCourse){
				return true
			}
		}
		state[course] = 2
		return false
	}

	for i:=0; i < numCourses; i++ {
		if state[i] == 0 {
			if hasCycle(i) {
				return false
			}
		}
	}
	return true
}

func main() {
	// Contoh 1: Aman
	fmt.Println(canFinish(2, [][]int{{1, 0}})) // Output: true

	// Contoh 2: Muter / Siklus (1 butuh 0, 0 butuh 1)
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}})) // Output: false
}