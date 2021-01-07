package main

import (
	"fmt"

	"github.com/davidcarbn/algorithms/data-structures/graph"
	"github.com/davidcarbn/algorithms/data-structures/graph/vertex"
	ga "github.com/davidcarbn/algorithms/graphalgoithms"
)

func main() {
	g := graph.New()

	v1 := *vertex.New(1)
	v2 := *vertex.New(2)
	v3 := *vertex.New(3)
	v4 := *vertex.New(4)
	v5 := *vertex.New(5)
	v6 := *vertex.New(6)
	v7 := *vertex.New(7)
	v8 := *vertex.New(8)

	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)
	g.AddVertex(v4)
	g.AddVertex(v5)
	g.AddVertex(v6)
	g.AddVertex(v7)
	g.AddVertex(v8)

	g.AddEdge(v1, v2)
	g.AddEdge(v2, v3)
	g.AddEdge(v3, v4)
	g.AddEdge(v3, v5)
	g.AddEdge(v4, v1)
	g.AddEdge(v5, v6)
	g.AddEdge(v6, v7)
	g.AddEdge(v7, v5)
	g.AddEdge(v7, v8)

	components := ga.ConnectedComponents(g)
	for i, comp := range components {
		fmt.Printf("%i: %v", i, comp)
	}
}
