package main

type SUVBuilder struct {
	make       string
	model      string
	engine     string
	wheels     int
	doors      int
	fuelType   string
	horsepower int
}

func newSUVBuilder() *SUVBuilder {
	return &SUVBuilder{}
}

func (b *SUVBuilder) setMake() {
	b.make = "BMW"
}

func (b *SUVBuilder) setModel() {
	b.model = "X5"
}

func (b *SUVBuilder) setEngine() {
	b.engine = "3.0L Turbocharged I6"
}

func (b *SUVBuilder) setWheels() {
	b.wheels = 4
}

func (b *SUVBuilder) setDoors() {
	b.doors = 5
}

func (b *SUVBuilder) setFuelType() {
	b.fuelType = "Hybrid"
}

func (b *SUVBuilder) setHorsepower() {
	b.horsepower = 335
}

func (b *SUVBuilder) getCar() Car {
	return Car{
		make:       b.make,
		model:      b.model,
		engine:     b.engine,
		wheels:     b.wheels,
		doors:      b.doors,
		fuelType:   b.fuelType,
		horsepower: b.horsepower,
	}
}