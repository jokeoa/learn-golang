package main

// IOTvisitor defines operations that can be performed on each device type.
type IOTvisitor interface {
	visitComputer(computer *Computer)

	visitServer(server *Server)

	visitPhone(phone *Phone)
	visitPrinter(printer *Printer)
}
