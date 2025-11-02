package main

import "fmt"

// Delivery is the subject that observers watch.
// When its status changes, all registered observers get notified.
type Delivery struct {
	observerList []Observer
	trackingID   string
	status       string
	location     string
}

func newDelivery(trackingID string) *Delivery {
	return &Delivery{
		trackingID: trackingID,
		status:     "processing",
	}
}

// updateStatus changes the delivery state and tells everyone watching.
// You call this when the package moves or its status changes.
func (d *Delivery) updateStatus(status, location string) {
	d.status = status
	d.location = location
	fmt.Printf("Delivery %s: %s at %s\n", d.trackingID, status, location)
	d.notifyAll()
}

func (d *Delivery) register(o Observer) {
	d.observerList = append(d.observerList, o)
}

func (d *Delivery) deregister(o Observer) {
	d.observerList = removeFromslice(d.observerList, o)
}

// notifyAll loops through observers and updates each one.
// This is where the pattern shines - one change, many notifications.
func (d *Delivery) notifyAll() {
	message := fmt.Sprintf("Status: %s | Location: %s", d.status, d.location)
	for _, observer := range d.observerList {
		observer.update(message)
	}
}