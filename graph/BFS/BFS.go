package main

import "fmt"

type Graph struct {
	vertices map[string][]string
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

func (g *Graph) AddEdge(v1, v2 string) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

func (g *Graph) BFS(startVertex string) {
	visited := make(map[string] bool)
	var queue []string

	visited[startVertex] = true
	queue = append(queue, startVertex)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		neighbors := g.vertices[curr]

		fmt.Printf("%s -> ", curr)
		for _, neighbor := range neighbors {
			if !visited[neighbor]{
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println("end")
}


func main () {
	network := NewGraph()

	network.AddEdge("Andi", "Budi")
	network.AddEdge("Andi", "Cici")
	network.AddEdge("Budi", "Dedi")
	network.AddEdge("Cici", "Edo")
	network.AddEdge("Dedi", "Edo")

	network.BFS("Andi")

}