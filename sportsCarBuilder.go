package main

type SportsCarBuilder struct {
	make       string
	model      string
	engine     string
	wheels     int
	doors      int
	fuelType   string
	horsepower int
}

func newSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{}
}

func (b *SportsCarBuilder) setMake() {
	b.make = "Ferrari"
}

func (b *SportsCarBuilder) setModel() {
	b.model = "488 GTB"
}

func (b *SportsCarBuilder) setEngine() {
	b.engine = "3.9L Twin Turbo V8"
}

func (b *SportsCarBuilder) setWheels() {
	b.wheels = 4
}

func (b *SportsCarBuilder) setDoors() {
	b.doors = 2
}

func (b *SportsCarBuilder) setFuelType() {
	b.fuelType = "Gasoline"
}

func (b *SportsCarBuilder) setHorsepower() {
	b.horsepower = 661
}

func (b *SportsCarBuilder) getCar() Car {
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