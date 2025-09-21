package main

type ICar interface {
	GetBrand() string
	GetModel() string
	GetMaxSpeed() int
	GetEngineType() string
	Start() string
	Accelerate() string
	Stop() string
}

type Car struct {
	brand      string
	model      string
	maxSpeed   int
	engineType string
}

func (c *Car) GetBrand() string {
	return c.brand
}

func (c *Car) GetModel() string {
	return c.model
}

func (c *Car) GetMaxSpeed() int {
	return c.maxSpeed
}

func (c *Car) GetEngineType() string {
	return c.engineType
}

func (c *Car) Start() string {
	return "Engine started"
}

func (c *Car) Stop() string {
	return "Engine stopped"
}

func (c *Car) Accelerate() string {
	return "Accelerating"
}
