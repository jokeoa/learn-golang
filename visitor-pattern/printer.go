package main

// Printer represents a printer device.
type Printer struct {
	Brand string
	Model string
	Price float64
}

// Accept lets a visitor operate on the Printer.
func (p *Printer) Accept(visitor IOTvisitor) {
	visitor.visitPrinter(p)
}
