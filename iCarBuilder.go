package main

type ICarBuilder interface {
	setMake()
	setModel()
	setEngine()
	setWheels()
	setDoors()
	setFuelType()
	setHorsepower()
	getCar() Car
}

func getCarBuilder(builderType string) ICarBuilder {
	if builderType == "sports" {
		return newSportsCarBuilder()
	}

	if builderType == "suv" {
		return newSUVBuilder()
	}
	return nil
}