package main

func main() {
	connector := &IOTconnector{}

	computer := &Computer{Brand: "Dell", Model: "XPS 15", Price: 1999.99}
	server := &Server{Brand: "HP", Model: "ProLiant DL380", Price: 4999.00}
	phone := &Phone{Brand: "Apple", Model: "iPhone 15", Price: 1299.00}
	printer := &Printer{Brand: "Canon", Model: "LBP223dw", Price: 249.99}

	computer.Accept(connector)
	server.Accept(connector)
	phone.Accept(connector)
	printer.Accept(connector)
}
