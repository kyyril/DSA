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

func (g *Graph) DFS(vertex string) {
	visited := make(map[string] bool)
	g.helperDfs(vertex, visited)
}

func (g *Graph) helperDfs(vertex string, visited map[string] bool) {
	visited[vertex] = true //inisiasi vertex/neighbors
	fmt.Printf("%s -> ", vertex)

	neighbors := g.vertices[vertex]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			g.helperDfs(neighbor, visited)
		}
	}
}

func (g *Graph) Display() {
	for vertex, neighbors := range g.vertices {
		fmt.Printf("%s -> %v\n", vertex, neighbors)
	}
}

func main () {
	network := NewGraph()

	network.AddEdge("Andi", "Budi")
	network.AddEdge("Andi", "Cici")
	network.AddEdge("Budi", "Dedi")
	network.AddEdge("Cici", "Edo")
	network.AddEdge("Dedi", "Edo") // Ini membuat siklus: Budi-Dedi-Edo-Cici-Andi
	
	network.Display()
	// Jalankan DFS dimulai dari "Andi"
	network.DFS("Andi")
}