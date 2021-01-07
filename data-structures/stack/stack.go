package stack

import (
	"errors"

	"github.com/davidcarbn/algorithms/data-structures/graph/vertex"
)

type Vertex = vertex.Vertex
type Stack struct {
	Elems []Vertex
}

func New() *Stack {
	return &Stack{}
}
func (s *Stack) IsEmpty() bool {
	return len(s.Elems) == 0
}
func (s *Stack) Push(value Vertex) {
	s.Elems = append(s.Elems, value)
}
func (s *Stack) Peek() (Vertex, error) {
	var top Vertex
	if s.IsEmpty() {
		return top, errors.New("Stack.Pop() Error: Stack is Empty")
	}
	return s.Elems[len(s.Elems)-1], nil
}
func (s *Stack) Pop() (Vertex, error) {
	var top Vertex
	if s.IsEmpty() {
		return top, errors.New("Stack.Pop() Error: Stack is Empty")
	}
	top = s.Elems[len(s.Elems)-1]
	s.Elems = s.Elems[:len(s.Elems)-1]
	return top, nil
}
