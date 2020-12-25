package graph

import (
	"fmt"
	"testing"

	"github.com/davidcarbn/algorithms/data-structures/graph/vertex"
)

type Vertex = vertex.Vertex

type Graph struct {
	vertices []*Vertex
	edges    map[*Vertex][]*Vertex
}

func New() *Graph {
	return new(Graph)
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

func TestAddVertex(t *testing.T) {
	var g Graph
	v1 := Vertex{1}
	v2 := Vertex{2}
	g.AddVertex(&v1)
	g.AddVertex(&v2)
	for _, v := range g.vertices {
		if v != &v1 || v != &v2 {
			t.Errorf("AddVertex = %q", g.vertices)
		}
	}
}
