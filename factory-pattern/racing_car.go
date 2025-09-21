package main

import "fmt"

// RacingCar represents a racing car implementation
// Following Single Responsibility Principle (SRP)
type RacingCar struct {
	Car
	turboEnabled bool
	aeroPackage  string
}

func NewRacingCar() ICar {
	return &RacingCar{
		Car: Car{
			brand:      "Ferrari",
			model:      "F1-2024",
			maxSpeed:   350,
			engineType: "V8 Turbo",
		},
		turboEnabled: true,
		aeroPackage:  "Advanced Aerodynamics",
	}
}

func (r *RacingCar) Accelerate() string {
	if r.turboEnabled {
		return fmt.Sprintf("Racing car accelerating with turbo boost! Max speed: %d km/h", r.maxSpeed)
	}
	return fmt.Sprintf("Racing car accelerating normally. Max speed: %d km/h", r.maxSpeed)
}

func (r *RacingCar) Start() string {
	return fmt.Sprintf("Racing engine roaring to life! %s ready for the track", r.model)
}

func (r *RacingCar) GetTurboStatus() bool {
	return r.turboEnabled
}

func (r *RacingCar) GetAeroPackage() string {
	return r.aeroPackage
}
