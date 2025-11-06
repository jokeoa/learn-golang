package main

// Computer represents a personal computer device.
type Computer struct {
	Brand string
	Model string
	Price float64
}

// Accept lets a visitor operate on the Computer.
func (c *Computer) Accept(visitor IOTvisitor) {
	visitor.visitComputer(c)
}
