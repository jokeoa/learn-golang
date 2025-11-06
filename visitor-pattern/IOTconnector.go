package main

import "fmt"

type IOTconnector struct {
	isConnected bool
}

func (c *IOTconnector) ConnectForComputer() {
}

func (c *IOTconnector) visitComputer(computer *Computer) {
	c.connect("Computer", computer.Brand, computer.Model)
}

func (c *IOTconnector) visitServer(server *Server) {
	c.connect("Server", server.Brand, server.Model)
}

func (c *IOTconnector) visitPhone(phone *Phone) {
	c.connect("Phone", phone.Brand, phone.Model)
}

func (c *IOTconnector) visitPrinter(printer *Printer) {
	c.connect("Printer", printer.Brand, printer.Model)
}

func (c *IOTconnector) connect(deviceType, brand, model string) {
	if !c.isConnected {
		c.isConnected = true
	}
	fmt.Printf("Connecting %s %s %s to local network... connected\n", deviceType, brand, model)
}
