package main

// Renderer is the implementation interface
type Renderer interface {
	RenderCircle(radius int)
	RenderSquare(size int)
}