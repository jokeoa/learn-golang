package main

import "fmt"

func main() {
	sportsBuilder := getCarBuilder("sports")
	suvBuilder := getCarBuilder("suv")

	carDirector := newCarDirector(sportsBuilder)
	sportsCar := carDirector.buildCar()

	fmt.Printf("Sports Car Make: %s\n", sportsCar.make)
	fmt.Printf("Sports Car Model: %s\n", sportsCar.model)
	fmt.Printf("Sports Car Engine: %s\n", sportsCar.engine)
	fmt.Printf("Sports Car Horsepower: %d\n", sportsCar.horsepower)
	fmt.Printf("Sports Car Doors: %d\n", sportsCar.doors)
	fmt.Printf("Sports Car Fuel Type: %s\n", sportsCar.fuelType)

	carDirector.setBuilder(suvBuilder)
	suvCar := carDirector.buildCar()

	fmt.Printf("\nSUV Make: %s\n", suvCar.make)
	fmt.Printf("SUV Model: %s\n", suvCar.model)
	fmt.Printf("SUV Engine: %s\n", suvCar.engine)
	fmt.Printf("SUV Horsepower: %d\n", suvCar.horsepower)
	fmt.Printf("SUV Doors: %d\n", suvCar.doors)
	fmt.Printf("SUV Fuel Type: %s\n", suvCar.fuelType)
}
