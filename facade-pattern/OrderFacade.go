package main

import "fmt"

// OrderFacade hides the complexity of checking customer, payment, inventory, and notifications
type OrderFacade struct {
	customer     *Customer
	inventory    *Inventory
	payment      *Payment
	notification *Notification
	orderLog     *OrderLog
}

func newOrderFacade(customerID string, cardNumber string) *OrderFacade {
	fmt.Println("Setting up order system")
	orderFacade := &OrderFacade{
		customer:     newCustomer(customerID),
		payment:      newPayment(cardNumber),
		inventory:    newInventory(),
		notification: &Notification{},
		orderLog:     &OrderLog{},
	}
	fmt.Println("Order system ready")
	return orderFacade
}

func (o *OrderFacade) placeOrder(customerID string, cardNumber string, item string, quantity int) error {
	fmt.Printf("Processing order for %d x %s\n", quantity, item)

	// Check if customer exists
	err := o.customer.verifyCustomer(customerID)
	if err != nil {
		return err
	}

	// Verify payment method
	err = o.payment.verifyCard(cardNumber)
	if err != nil {
		return err
	}

	// Check if we have enough stock
	err = o.inventory.checkStock(item, quantity)
	if err != nil {
		return err
	}

	// Reduce stock
	o.inventory.reduceStock(item, quantity)

	// Send confirmation
	o.notification.sendOrderConfirmation(item, quantity)

	// Log the order
	o.orderLog.recordOrder(customerID, item, quantity)

	return nil
}