package main

import "fmt"

func runDeliveryExample() {
	fmt.Println("\n=== Delivery Observer Pattern ===\n")

	// Create a delivery to track
	delivery := newDelivery("DL-2024-001")

	// Set up observers who want updates
	customer := &RecipientCustomer{
		id:    "john_doe",
		phone: "+1-555-0123",
	}

	courier := &CourierApp{
		driverID: "driver_42",
	}

	warehouse := &WarehouseSystem{
		systemID: "WH-CENTRAL",
	}

	// Register them all
	delivery.register(customer)
	delivery.register(courier)
	delivery.register(warehouse)

	// Simulate the delivery journey
	// Each status change notifies all observers automatically
	delivery.updateStatus("dispatched", "Distribution Center")

	fmt.Println()
	delivery.updateStatus("in_transit", "Highway 101")

	fmt.Println()
	delivery.updateStatus("out_for_delivery", "Local Hub")

	// Customer no longer wants updates
	fmt.Println("\n[Customer unsubscribes from notifications]")
	delivery.deregister(customer)

	fmt.Println()
	delivery.updateStatus("delivered", "123 Main St")
}