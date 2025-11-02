package main

import "fmt"

// RecipientCustomer gets SMS notifications about their delivery.
// They care most about when it arrives.
type RecipientCustomer struct {
	id    string
	phone string
}

func (r *RecipientCustomer) update(message string) {
	fmt.Printf("→ SMS to %s (%s): Your delivery - %s\n", r.phone, r.id, message)
}

func (r *RecipientCustomer) getID() string {
	return r.id
}

// CourierApp updates the driver's mobile app.
// The driver needs real-time info to do their job.
type CourierApp struct {
	driverID string
}

func (c *CourierApp) update(message string) {
	fmt.Printf("→ App notification to driver %s: Package update - %s\n", c.driverID, message)
}

func (c *CourierApp) getID() string {
	return c.driverID
}

// WarehouseSystem logs everything for tracking and analytics.
// It's quiet but keeps records of every state change.
type WarehouseSystem struct {
	systemID string
}

func (w *WarehouseSystem) update(message string) {
	fmt.Printf("→ Warehouse log [%s]: %s\n", w.systemID, message)
}

func (w *WarehouseSystem) getID() string {
	return w.systemID
}