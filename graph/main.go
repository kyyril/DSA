package main 

import "fmt"

type Graph struct {
	vertices map[string] []string
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []string{}
	}
}

//undirected graph
func (g *Graph) AddEdge(v1, v2 string) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	//v1 -> v2; v2 -> v1
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

func (g *Graph) Display() {
	for vertex, neighbors := range g.vertices {
		fmt.Printf("%s -> %v\n", vertex, neighbors)
	}
}

func main () {
	network := NewGraph()
	network.AddVertex("jamal")
	network.AddVertex("wahyu")
	network.AddVertex("kiril")

	network.AddEdge("jamal", "wahyu")
	network.AddEdge("kiril", "wahyu")

	// network.AddEdge("kiril", "jamal")

	network.Display()
}