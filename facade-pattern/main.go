package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	orderSystem := newOrderFacade("customer123", "4532-1234-5678-9012")
	fmt.Println()

	err := orderSystem.placeOrder("customer123", "4532-1234-5678-9012", "Pasta", 2)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = orderSystem.placeOrder("customer123", "4532-1234-5678-9012", "Pizza", 1)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}