package main

import "fmt"

// Customer represents a registered customer
type Customer struct {
	id string
}

func newCustomer(customerID string) *Customer {
	return &Customer{
		id: customerID,
	}
}

func (c *Customer) verifyCustomer(customerID string) error {
	if c.id != customerID {
		return fmt.Errorf("customer ID doesn't match")
	}
	fmt.Println("Customer verified")
	return nil
}