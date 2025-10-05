package main

// Figure is the abstraction interface
type Figure interface {
	Draw()
	SetRenderer(Renderer)
}