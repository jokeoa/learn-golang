package main

type CarDirector struct {
	builder ICarBuilder
}

func newCarDirector(b ICarBuilder) *CarDirector {
	return &CarDirector{
		builder: b,
	}
}

func (d *CarDirector) setBuilder(b ICarBuilder) {
	d.builder = b
}

func (d *CarDirector) buildCar() Car {
	d.builder.setMake()
	d.builder.setModel()
	d.builder.setEngine()
	d.builder.setWheels()
	d.builder.setDoors()
	d.builder.setFuelType()
	d.builder.setHorsepower()
	return d.builder.getCar()
}
