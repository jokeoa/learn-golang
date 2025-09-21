package main

import (
	"fmt"
	"log"
)

func demonstrateCar(car ICar, carType string) {
	fmt.Printf("Brand: %s\n", car.GetBrand())
	fmt.Printf("Model: %s\n", car.GetModel())
	fmt.Printf("Engine: %s\n", car.GetEngineType())
	fmt.Printf("Max Speed: %d km/h\n", car.GetMaxSpeed())

	fmt.Printf("\n%s\n", car.Start())
	fmt.Printf("%s\n", car.Accelerate())
	fmt.Printf("%s\n", car.Stop())

	switch v := car.(type) {
	case *RacingCar:
		fmt.Printf("Turbo Enabled: %v\n", v.GetTurboStatus())
		fmt.Printf("Aero Package: %s\n", v.GetAeroPackage())
	case *SedanCar:
		fmt.Printf("Comfort Level: %s\n", v.GetComfortLevel())
		fmt.Printf("Fuel Efficiency: %.1f km/l\n", v.GetFuelEfficiency())
		fmt.Printf("Safety Rating: %d/5\n", v.GetSafetyRating())
	}
}

func handleCarCreation(factory CarFactory, carType CarType) {
	fmt.Printf("\nAttempting to create %s car...\n", carType)

	car, err := factory.CreateCar(carType)
	if err != nil {
		log.Printf("Error creating car: %v\n", err)
		return
	}

	demonstrateCar(car, string(carType))
}

func main() {

	factory := NewCarFactory()

	fmt.Printf("Supported car types: %v\n", factory.GetSupportedTypes())

	handleCarCreation(factory, RacingCarType)
	handleCarCreation(factory, SedanCarType)

	handleCarCreation(factory, "invalid")

	fmt.Printf("\nValidation Examples:\n")
	fmt.Printf("Is 'racing' a valid car type? %v\n", IsValidCarType("racing"))
	fmt.Printf("Is 'suv' a valid car type? %v\n", IsValidCarType("suv"))
}
