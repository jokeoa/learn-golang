package main

type IOTvisitor interface {
	visitComputer(computer *Computer)
	visitServer(server *Server)
	visitPhone(phone *Phone)
	visitPrinter(printer *Printer)
}
