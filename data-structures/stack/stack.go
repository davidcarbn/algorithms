package stack

type Stack struct {
	Elems []int
}

func (s *Stack) New() *Stack {
	return &Stack{}
}
func (s *Stack) Push(value int) {
	s.Elems = append(s.Elems,value)
}
func (s *Stack) Peek() int {
	return s.Elems[len(s.Elems)-1]
}
func (s *Stack) Pop() int {
	top := s:Elems[len(s.Elems)-1]
	s.Elems = s.Elems[:len(s.Elems)-1]
	return top
}