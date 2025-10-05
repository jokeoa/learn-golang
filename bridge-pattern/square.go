package main

// Square is a concrete abstraction
type Square struct {
	size     int
	renderer Renderer
}

func NewSquare(size int) *Square {
	return &Square{size: size}
}

func (s *Square) SetRenderer(r Renderer) {
	s.renderer = r
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.size)
}