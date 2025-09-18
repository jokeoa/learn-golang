# Software Design Patterns

## Assignment 1 (Builder Pattern)
    
The Builder pattern is a creational design pattern that provides a flexible solution to construct complex objects step by step. This implementation includes two examples:


### Car Builder Example
- **Car**: Product with `make`, `model`, `engine`, `wheels`, `doors`, `fuelType`, and `horsepower` properties
- **SportsCarBuilder**: Creates high-performance sports cars (Ferrari 488 GTB)
- **SUVBuilder**: Creates family SUVs (BMW X5)
- **CarDirector**: Manages the car construction process

### UML Diagram
![Builder Pattern UML](https://github.com/jokeoa/learn-golang/blob/main/images/builderDiagram.png?raw=true)

### How to Run

```bash
# Build the project
go build .

# Run the demonstration
go run .

# Format code
go fmt .

# Run tests
go test .
```