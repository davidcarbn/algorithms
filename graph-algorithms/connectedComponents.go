package connectedcomponents

import (
	"github.com/davidcarbn/algorithms/data-structures/graph"
	"github.com/davidcarbn/algorithms/data-structures/stack"
)

type Graph = graph.Graph
type Vertex = graph.Vertex
type Stack = stack.Stack

var dfsnum map[Vertex]int
var roots *Stack
var unfinished *Stack
var isUnfinished map[Vertex]bool
var visited map[Vertex]bool
var count int
var connectedComponents [][]Vertex

func ConnectedComponents(g *Graph) [][]Vertex {
	vertices := g.GetVertices()
	dfsnum = make(map[Vertex]int)
	roots = stack.New()
	unfinished = stack.New()
	isUnfinished = make(map[Vertex]bool)
	visited = make(map[Vertex]bool)
	count = 0
	connectedComponents := make([][]Vertex, 0)
	dfs(vertices[0], g)
	return connectedComponents
}

func dfs(v Vertex, g *Graph) {
	count++
	dfsnum[v] = count
	visited[v] = true
	for _, w := range g.GetOutEdges(v) {
		if _, ok := visited[w]; !ok {
			dfs(w, g)
		} else if _, ok := isUnfinished[w]; ok {
			topRoot, err := roots.Peek()
			if err == nil {
				for valW, valRoot := dfsnum[w], dfsnum[topRoot]; valW < valRoot; {
					roots.Pop()
					topRoot, _ = roots.Peek()
				}
			}
		}
	}
	if topRoot, _ := roots.Peek(); topRoot == v {
		comp := make([]Vertex, 0)
		for {
			w, _ := unfinished.Pop()
			isUnfinished[w] = false
			comp = append(comp, w)
			if w == v {
				break
			}
		}
		roots.Pop()
		connectedComponents = append(connectedComponents, comp)
	}
}
