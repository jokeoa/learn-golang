package main

type Printer struct {
	Brand string
	Model string
	Price float64
}

func (p *Printer) Accept(visitor IOTvisitor) {
	visitor.visitPrinter(p)
}
