package main

import "fmt"

// DemoFigureRendering demonstrates the bridge pattern with figures and renderers
func DemoFigureRendering() {

	// Create renderers (implementations)
	asciiRenderer := &AsciiRenderer{}
	guiRenderer := &GuiRenderer{}

	// Create figures (abstractions)
	circle := NewCircle(5)
	square := NewSquare(10)

	// Circle with ASCII renderer
	circle.SetRenderer(asciiRenderer)
	circle.Draw()
	fmt.Println()

	// Circle with GUI renderer
	circle.SetRenderer(guiRenderer)
	circle.Draw()
	fmt.Println()

	// Square with ASCII renderer
	square.SetRenderer(asciiRenderer)
	square.Draw()
	fmt.Println()

	// Square with GUI renderer
	square.SetRenderer(guiRenderer)
	square.Draw()
	fmt.Println()
}
