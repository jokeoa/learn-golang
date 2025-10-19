# Facade Pattern: Order System

## What This Does

This is a restaurant order system built in Go. It shows how the Facade pattern hides complexity behind a simple interface.

![UML Class Diagram](https://uml.planttext.com/plantuml/png/TLDDJyCm3BttL-JOibNYd0CQ6aAQX0RIn0bno58FeZL9STAXWlZlP4tJZspjAV5p_5wVxPMn36qjoyGSXQeOCabuoOovjnR5lhQmuLhjn0-GsoHf4R8k1IXBVXFYleLaGgLPyXVGPom07nbd1NGfEQ-DrHAGadLtyd6XJgwGneQIJJXwf6ADT7TANlp1HvMsuY0uiq8hIhQZcyVrcVf9Vr0lmfsyH0NV8rLfaBDvM1BEiDZMSW-OqR5q_x0g6OSsElDeIOG5kIHVDLDMsCOL0KIDWntXZuCNeg34E6lLe2saZ3YooAh52L0SckkP4MDdMMz-B_Q2kCjC9oQC7G-q3fycHPHzzmAfKyUFb4XMlNdETn5AEju-FStYpY_u9_3ZxlmydR6sJKCeQWvNymPPuma9usf05Mr7rreT1CeMd2yH1YayH-1koYv7uag1Rato_6wo4vIy6aZ6aPiy7xeT0w4TSNomBKOcUn63_M8aAsU4s_j_)

## Structure

**Main Components:**

- **Client** - Your application code that places orders
- **OrderFacade** - The simplified interface that orchestrates everything
- **Five Subsystems** - Customer, Payment, Inventory, Notification, OrderLog


**Files:**

- `main.go` - Entry point and client code
- `OrderFacade.go` - The facade implementation
- `customer.go` - Customer verification
- `payment.go` - Payment processing
- `inventory.go` - Stock management with initial data (Pizza: 10, Pasta: 15, Salad: 20)
- `notification.go` - Order confirmations and logging

## Running The Code

```bash
go run .
```

You'll see detailed output showing each step of the order process for two sample orders.