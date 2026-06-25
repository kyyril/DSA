package main
import "fmt"


type DirectedGraph struct {
	vertices map[string] []string
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{vertices: make(map[string][]string)}
} 
	
func(g *DirectedGraph) AddDirectedEdge(v1, v2 string){
	g.vertices[v1] = append(g.vertices[v1], v2)
}
	
//DFS 3-State
func(g *DirectedGraph) HasCycle() bool{
	//record state each of point
	//0:never; 1:current; 2:complate
	state := make(map[string] int)
		
	for vertex := range g.vertices {
		if state[vertex] == 0 {
			if g.dfsCheckCycle(vertex, state){
				return true
			}
		}
	}
	return false
}

func(g *DirectedGraph) dfsCheckCycle(vertex string, state map[string] int) bool{
	state[vertex] = 1

	neighbors := g.vertices[vertex] 
	for _, neighbor := range neighbors {
		if state[neighbor] == 1 {
			return true
		}

		if state[neighbor] == 0 {
			if g.dfsCheckCycle(neighbor, state){
				return true
			}
		}
	}
	state[vertex] = 2
	return false
}

func main () {
	saveGraph := NewDirectedGraph()
	saveGraph.AddDirectedEdge("A", "B")
	saveGraph.AddDirectedEdge("B", "C")
	fmt.Printf("Graph have a cycle? %t\n", saveGraph.HasCycle()) // Output: false
	
	unsaveGraph := NewDirectedGraph()
	unsaveGraph.AddDirectedEdge("X", "Y")
	unsaveGraph.AddDirectedEdge("Y", "Z")
	unsaveGraph.AddDirectedEdge("Z", "X")
	fmt.Printf("Graph have a cycle? %t\n", unsaveGraph.HasCycle()) // Output: true
}
// Untuk Graph yang memiliki arah panah (Directed Graph), kita tidak bisa hanya menggunakan map visited biasa. Kita butuh 3 status untuk setiap titik:

// Belum Dikunjungi (0): Titik yang sama sekali belum disentuh.
// Sedang Dikunjungi (1): Titik yang sedang masuk dalam jalur pencarian DFS saat ini (masih ada di dalam call stack).
// Selesai Dikunjungi (2): Titik yang sudah selesai dijelajahi beserta seluruh anak-cucunya, dan terbukti aman (tidak ada siklus).

// Kunci Utama: Jika saat melakukan DFS kita bertemu dengan titik yang statusnya "Sedang Dikunjungi (1)", artinya kita mendeteksi adanya Siklus! Kita menabrak punggung kita sendiri saat sedang berjalan berputar.