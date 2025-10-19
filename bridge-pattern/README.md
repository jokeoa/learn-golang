# Bridge Pattern
[![](https://mermaid.ink/img/pako:eNq1VV2L4jAU_SshMKBdLVqtToMIg7LLvszC-Lb0JdvGGqZN3Xzg7rj-902TVFu1sAszgpqenHvuyb03eoRJmRKIYJJjIdYUZxwXMQP69fAAnn4IyXEiacnAhqbEbhgq-EwzxQk4Wqx6LRaUScK3OCHL5QX-tOb40Os3gA2RL4SlhBPeqxdu_2S_molWlCd5K9GQ45QqgYDO10SdFAK1aCPnMzlYoZ4NrmL7wLNYhzcOrtzdHOeO3c1Phdt1GQr6Rv7XrJXpVaHOqkXex6ru7ddin5OCMInvtreW-5cGW-5teW8o14fqLOKTSCi9Z-EDcn1RH59J13uVU1JPgBtsg7RyFpiyroa9kNy0SuzoXljU3cHFH9-vr4mes7qxd0huOG9J9vNcB0NuN-E2ps1ulrFL33kcDpegwVWCOJaz17nvVGzdKpY7GwIJJ1jWtAahIdOinKv6XGoMbEsOtnl5sDDTmIGcUAzHur423J2hF_bjmAUXuFUsvTXRWzjPBUgM329eVlxxq_jpFcndV9j0eDFjixfDNclJZpKmmk9ZppVkCagUoP5V6ZBodzSG35TcK4liQ_M8r1p4ZhKrpeem8gp2VAgHMOM0hUhyRQawIFxPr36EZqBjKHd6AGKI9DLF_LWydNIxe8y-l2VRh_FSZTuItjgX-kntU30w9zd0Ru2pVqViEqJxOJsbFYiO8BdEw_Fj4IfhKAyiYBpG8ygIBvC3xsPQj4JJoOH5ZBbq92kA30zmwB_PotHocRpNpsFkOgtOfwH5oCGV?type=png)](https://mermaid.live/edit#pako:eNq1VV2L4jAU_SshMKBdLVqtToMIg7LLvszC-Lb0JdvGGqZN3Xzg7rj-902TVFu1sAszgpqenHvuyb03eoRJmRKIYJJjIdYUZxwXMQP69fAAnn4IyXEiacnAhqbEbhgq-EwzxQk4Wqx6LRaUScK3OCHL5QX-tOb40Os3gA2RL4SlhBPeqxdu_2S_molWlCd5K9GQ45QqgYDO10SdFAK1aCPnMzlYoZ4NrmL7wLNYhzcOrtzdHOeO3c1Phdt1GQr6Rv7XrJXpVaHOqkXex6ru7ddin5OCMInvtreW-5cGW-5teW8o14fqLOKTSCi9Z-EDcn1RH59J13uVU1JPgBtsg7RyFpiyroa9kNy0SuzoXljU3cHFH9-vr4mes7qxd0huOG9J9vNcB0NuN-E2ps1ulrFL33kcDpegwVWCOJaz17nvVGzdKpY7GwIJJ1jWtAahIdOinKv6XGoMbEsOtnl5sDDTmIGcUAzHur423J2hF_bjmAUXuFUsvTXRWzjPBUgM329eVlxxq_jpFcndV9j0eDFjixfDNclJZpKmmk9ZppVkCagUoP5V6ZBodzSG35TcK4liQ_M8r1p4ZhKrpeem8gp2VAgHMOM0hUhyRQawIFxPr36EZqBjKHd6AGKI9DLF_LWydNIxe8y-l2VRh_FSZTuItjgX-kntU30w9zd0Ru2pVqViEqJxOJsbFYiO8BdEw_Fj4IfhKAyiYBpG8ygIBvC3xsPQj4JJoOH5ZBbq92kA30zmwB_PotHocRpNpsFkOgtOfwH5oCGV)
A Go implementation showing how to separate what you draw from how you draw it.

## The Problem

You have shapes (Circle, Square) and ways to render them (ASCII, GUI). Without the Bridge pattern, you'd need a class for every combination:
- AsciiCircle
- GuiCircle
- AsciiSquare
- GuiSquare

That's 4 classes for just 2 shapes and 2 renderers. Add a Triangle and an SVG renderer? Now you need 9 classes.

## The Solution

Split things into two parts:
1. **Shapes** - Circle, Square
2. **Renderers** - ASCII, GUI

Each shape holds a reference to a renderer. You can swap renderers at runtime.

## Files

- `figure.go` - Shape interface
- `circle.go`, `square.go` - Concrete shapes
- `renderer.go` - Renderer interface
- `ascii_renderer.go`, `gui_renderer.go` - Concrete renderers
- `main.go` - Demo

## How to Use

```go
// Make a circle
circle := NewCircle(5)

// Draw it with ASCII
circle.SetRenderer(&AsciiRenderer{})
circle.Draw()

// Switch to GUI
circle.SetRenderer(&GuiRenderer{})
circle.Draw()
```

## Run It

```bash
go run .
```

You'll see the same circle drawn two different ways.

## Add a New Shape

```go
type Triangle struct {
    base     int
    renderer Renderer
}

func (t *Triangle) Draw() {
    t.renderer.RenderTriangle(t.base)
}
```

You also need to add `RenderTriangle()` to the Renderer interface and implement it in each renderer.

## Add a New Renderer

```go
type SvgRenderer struct{}

func (s *SvgRenderer) RenderCircle(radius int) {
    fmt.Printf("<circle r='%d' />\n", radius)
}

func (s *SvgRenderer) RenderSquare(size int) {
    fmt.Printf("<rect width='%d' height='%d' />\n", size, size)
}
```

## Why This Works

Adding a new shape? One class.
Adding a new renderer? One class.

No matter how many shapes and renderers you have, you never need a class for every combination.

## When to Use This

- You have two things that vary independently
- You want to change behavior at runtime
- You're creating too many classes because of combinations