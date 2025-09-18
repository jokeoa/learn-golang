# Software Design Patterns

This project demonstrates various design patterns in Go using a workspace structure with separate modules for each pattern.

## Project Structure

```
learn/
├── go.work                 # Go workspace file
├── builder-pattern/        # Builder pattern implementation
│   ├── go.mod
│   ├── main.go
│   └── ...
├── factory-pattern/        # Factory pattern implementation
│   ├── go.mod
│   ├── main.go
│   └── ...
└── images/                 # UML diagrams
```

## Assignment 1 (Builder Pattern)

The Builder pattern is a creational design pattern that provides a flexible solution to construct complex objects step by step.

### Car Builder Example
- **Car**: Product with `make`, `model`, `engine`, `wheels`, `doors`, `fuelType`, and `horsepower` properties
- **SportsCarBuilder**: Creates high-performance sports cars (Ferrari 488 GTB)
- **SUVBuilder**: Creates family SUVs (BMW X5)
- **CarDirector**: Manages the car construction process

### UML Diagram
![Builder Pattern UML](https://github.com/jokeoa/learn-golang/blob/main/images/builderDiagram.png?raw=true)

## How to Run

### Run a specific pattern:

```bash
# Run builder pattern
cd builder-pattern
go run .

# Run factory pattern
cd factory-pattern
go run .
```

### Run from workspace root:

```bash
# Run builder pattern from root
go run ./builder-pattern

# Run factory pattern from root
go run ./factory-pattern
```

### Other commands:

```bash
# Format all code in workspace
go fmt ./...

# Build all modules
go work sync

# Run tests for all modules
go test ./...
```