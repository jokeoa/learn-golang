package main

import "fmt"

type SedanCar struct {
	Car
	comfortLevel   string
	fuelEfficiency float64
	safetyRating   int
}

func NewSedanCar() ICar {
	return &SedanCar{
		Car: Car{
			brand:      "Toyota",
			model:      "Camry",
			maxSpeed:   180,
			engineType: "Hybrid",
		},
		comfortLevel:   "Premium",
		fuelEfficiency: 15.5, // km per liter
		safetyRating:   5,
	}
}

func (s *SedanCar) Accelerate() string {
	return fmt.Sprintf("Sedan accelerating smoothly and efficiently. Max speed: %d km/h", s.maxSpeed)
}

func (s *SedanCar) Start() string {
	return fmt.Sprintf("Sedan engine starting quietly. %s ready for comfortable driving", s.model)
}

func (s *SedanCar) GetComfortLevel() string {
	return s.comfortLevel
}

func (s *SedanCar) GetFuelEfficiency() float64 {
	return s.fuelEfficiency
}

func (s *SedanCar) GetSafetyRating() int {
	return s.safetyRating
}
