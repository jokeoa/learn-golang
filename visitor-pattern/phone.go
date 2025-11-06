package main

// Phone represents a mobile phone device.
type Phone struct {
	Brand string
	Model string
	Price float64
}

// Accept lets a visitor operate on the Phone.
func (p *Phone) Accept(visitor IOTvisitor) {
	visitor.visitPhone(p)
}
