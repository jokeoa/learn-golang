package main

import "fmt"

func main() {
	// Original example: item stock notification
	fmt.Println("=== Item Stock Observer Pattern ===\n")

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()

	// New example: delivery tracking
	runDeliveryExample()
}