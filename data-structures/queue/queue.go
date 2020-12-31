package queue

import (
	"github.com/davidcarbn/algorithms/data-structures/graph/vertex"
)

type Vertex = vertex.Vertex

type Queue struct {
	Items []Vertex
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v Vertex) {
	q.Items = append(q.Items, v)
}
func (q *Queue) Pop() Vertex {
	item := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return item
}
