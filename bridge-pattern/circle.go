package main

// Circle is a concrete abstraction
type Circle struct {
	radius   int
	renderer Renderer
}

func NewCircle(radius int) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) SetRenderer(r Renderer) {
	c.renderer = r
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}