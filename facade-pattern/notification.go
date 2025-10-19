package main

import "fmt"

// Notification sends order confirmations
type Notification struct {
}

func (n *Notification) sendOrderConfirmation(item string, quantity int) {
	fmt.Printf("Order confirmed: %d x %s\n", quantity, item)
}

// OrderLog keeps track of all orders
type OrderLog struct {
}

func (l *OrderLog) recordOrder(customerID string, item string, quantity int) {
	fmt.Printf("Logged: Customer %s ordered %d x %s\n", customerID, quantity, item)
}