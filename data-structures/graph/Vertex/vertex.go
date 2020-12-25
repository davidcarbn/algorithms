package vertex

type Vertex struct {
	Elem int
}

func New(i int) *Vertex {
	return &Vertex{i}
}
