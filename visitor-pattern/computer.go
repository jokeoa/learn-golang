package main

type Computer struct {
	Brand string
	Model string
	Price float64
}

func (c *Computer) Accept(visitor IOTvisitor) {
	visitor.visitComputer(c)
}
