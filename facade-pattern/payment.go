package main

import "fmt"

// Payment handles credit card verification
type Payment struct {
	cardNumber string
}

func newPayment(cardNumber string) *Payment {
	return &Payment{
		cardNumber: cardNumber,
	}
}

func (p *Payment) verifyCard(cardNumber string) error {
	if p.cardNumber != cardNumber {
		return fmt.Errorf("card number doesn't match")
	}
	fmt.Println("Payment verified")
	return nil
}