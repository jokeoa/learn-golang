package main

type Phone struct {
	Brand string
	Model string
	Price float64
}

func (p *Phone) Accept(visitor IOTvisitor) {
	visitor.visitPhone(p)
}
