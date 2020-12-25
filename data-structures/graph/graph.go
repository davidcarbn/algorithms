package graph

import "fmt"

type Vertex struct {
	elem int
}

type Graph struct {
	vertices []*Vertex
	edges    map[*Vertex][]*Vertex
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}
func (g *Graph) AddEdge(v *Vertex, w *Vertex) {
	if g.edges == nil {
		g.edges = make(map[*Vertex][]*Vertex)
	}
	g.edges[v] = append(g.edges[v], w)
	fmt.Printf("%p", g.edges)
}

func main() {
	var g Graph
	v1 := Vertex{1}
	v2 := Vertex{2}
	v3 := Vertex{3}
	v4 := Vertex{3}

	g.AddVertex(&v1)
	g.AddVertex(&v2)
	g.AddVertex(&v3)
	g.AddVertex(&v4)

	g.AddEdge(&v1, &v2)
	g.AddEdge(&v2, &v3)
	g.AddEdge(&v2, &v4)
	g.AddEdge(&v3, &v4)
}
