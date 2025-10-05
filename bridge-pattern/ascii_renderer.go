package main

import "fmt"

// AsciiRenderer renders figures in ASCII format
type AsciiRenderer struct{}

func (a *AsciiRenderer) RenderCircle(radius int) {
	fmt.Printf("ASCII Circle with radius %d:\n", radius)
	fmt.Println("   ***   ")
	fmt.Println(" *     * ")
	fmt.Println("*       *")
	fmt.Println(" *     * ")
	fmt.Println("   ***   ")
}

func (a *AsciiRenderer) RenderSquare(size int) {
	fmt.Printf("ASCII Square with size %d:\n", size)
	fmt.Println("*********")
	fmt.Println("*       *")
	fmt.Println("*       *")
	fmt.Println("*       *")
	fmt.Println("*********")
}