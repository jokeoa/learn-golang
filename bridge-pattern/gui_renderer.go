package main

import "fmt"

// GuiRenderer renders figures in GUI format
type GuiRenderer struct{}

func (g *GuiRenderer) RenderCircle(radius int) {
	fmt.Printf("GUI Circle with radius %d:\n", radius)
	fmt.Println("[Opening GUI window...]")
	fmt.Printf("[Drawing circle using graphics library with radius=%d]\n", radius)
	fmt.Println("[Applying anti-aliasing and smooth edges]")
}

func (g *GuiRenderer) RenderSquare(size int) {
	fmt.Printf("GUI Square with size %d:\n", size)
	fmt.Println("[Opening GUI window...]")
	fmt.Printf("[Drawing square using graphics library with size=%d]\n", size)
	fmt.Println("[Applying shadows and 3D effects]")
}