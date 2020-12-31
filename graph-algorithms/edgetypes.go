package edgetypes

import (
	"fmt"

	"github.com/davidcarbn/algorithms/data-structures/graph"
	"github.com/davidcarbn/algorithms/data-structures/graph/vertex"
)

type Vertex = vertex.Vertex
type Graph = graph.Graph

func getEdgeTypes(g *Graph) map[string][]Vertex {
	var edgeTypes = make(map[string][]Vertex)
	for v := range g.GetVertices() {
		fmt.Println(v)
	}

	return edgeTypes
}
