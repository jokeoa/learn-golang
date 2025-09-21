package main

import (
	"fmt"
	"strings"
)

type CarType string

const (
	RacingCarType CarType = "racing"
	SedanCarType  CarType = "sedan"
)

type CarFactory interface {
	CreateCar(carType CarType) (ICar, error)
	GetSupportedTypes() []CarType
}

type ConcreteCarFactory struct{}

func NewCarFactory() CarFactory {
	return &ConcreteCarFactory{}
}

func (f *ConcreteCarFactory) CreateCar(carType CarType) (ICar, error) {
	switch strings.ToLower(string(carType)) {
	case string(RacingCarType):
		return NewRacingCar(), nil
	case string(SedanCarType):
		return NewSedanCar(), nil
	default:
		return nil, fmt.Errorf("unsupported car type: %s. Supported types: %v",
			carType, f.GetSupportedTypes())
	}
}

func (f *ConcreteCarFactory) GetSupportedTypes() []CarType {
	return []CarType{RacingCarType, SedanCarType}
}

func IsValidCarType(carType CarType) bool {
	factory := NewCarFactory()
	supportedTypes := factory.GetSupportedTypes()

	for _, supportedType := range supportedTypes {
		if strings.EqualFold(string(carType), string(supportedType)) {
			return true
		}
	}
	return false
}
