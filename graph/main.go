package main
import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key int
	adjacent []*Vertex
}

// graph represent an adjacency list graph


// vertex represent an adjacency matrix graph

// add vertex
func (g *Graph) AddVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key:k})
}
// add edge

func main(){
	test := &Graph{}
	for i:=1; i <= 5; i++ {
		test.AddVertex(i)
	}
	fmt.Println(test)
}