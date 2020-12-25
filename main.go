package main

import (
	"github.com/davidcarbn/algorithms/data-structures/graph"
	"github.com/davidcarbn/algorithms/data-structures/graph/vertex"
)

func main() {
	g := graph.New()

	v1 := vertex.New(1)
	v2 := vertex.New(2)
	v3 := vertex.New(3)
	v4 := vertex.New(4)

	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)
	g.AddVertex(v4)

	g.AddEdge(v1, v2)
	g.AddEdge(v1, v3)
	g.AddEdge(v2, v3)
	g.AddEdge(v2, v4)
}
