// Vertex in a Graph with unique ID
package vertex

type Vertex struct {
	Id int
}

func New(i int) *Vertex {
	return &Vertex{i}
}
func (v *Vertex) GetId() int {
	return v.Id
}
