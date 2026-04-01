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
	if contains(g.vertices, k){
		err := fmt.Errorf("\nexisting key %v", k)
		fmt.Println(err.Error())
	}else {
		g.vertices = append(g.vertices, &Vertex{key:k})
	}
}
// add edge
func (g *Graph) AddEdge(from, to int){
	//get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge %v -> %v", from, to)
		fmt.Println(err.Error())
	}else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex) 
	}
}

//get vertex
func (g *Graph) GetVertex(k int) *Vertex{
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//check contains
func contains(s []*Vertex, k int) bool{
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
//print all
func(g *Graph) Print(){
	for _, v := range g.vertices {
		fmt.Printf("\nvertex: %v", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("-> %v", v.key)
		}
	}
}

func main(){
	test := &Graph{}
	for i:=1; i <= 5; i++ {
		test.AddVertex(i)
	}
	test.AddVertex(1)
	test.AddEdge(0,0)
	test.AddEdge(1,2)
	test.Print()
}