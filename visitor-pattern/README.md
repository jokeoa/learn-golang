## Visitor Pattern: IoT Devices (Go)

This project shows the Visitor pattern in Go using simple IoT devices.

### What it does
- **Devices**: `Computer`, `Server`, `Phone`, `Printer`
- **Visitor**: `IOTconnector` connects any device to a local network
- **Contract**: `IOTvisitor` defines visit methods per device type

You add behavior to many types without changing their code. You write a visitor and let each device accept it.

### Why use Visitor
- Add new operations without touching device structs
- Keep device data separate from behaviors
- Avoid big `switch`/`type` checks spread across code

### Project structure
- `main.go`: Builds devices and applies the visitor
- `IOTvisitor.go`: Visitor interface
- `IOTconnector.go`: Visitor implementation (connects devices)
- `computer.go`, `server.go`, `phone.go`, `printer.go`: Device structs with `Accept`
- `go.mod`: Module file

### How it works
Each device has an `Accept(visitor IOTvisitor)` method. It calls the matching `visitX` method on the visitor.

Flow in `main.go`:
1. Create `IOTconnector`
2. Create devices
3. Call `Accept` on each device with the visitor

`IOTconnector` keeps a simple `isConnected` flag and prints connection messages.

### Run it
```bash
go run .
```

Expected output (similar):
```text
Connecting Computer Dell XPS 15 to local network... connected
Connecting Server HP ProLiant DL380 to local network... connected
Connecting Phone Apple iPhone 15 to local network... connected
Connecting Printer Canon LBP223dw to local network... connected
```

### Extend it
Add a new device:
1. Create a struct with fields you need
2. Implement `Accept(visitor IOTvisitor)` and call the right `visitX`
3. Add a `visitNewDevice(*NewDevice)` method to `IOTvisitor`
4. Implement that method on `IOTconnector` (or any visitor)

Add a new behavior:
1. Create a new visitor type (e.g., `IOTAuditor`)
2. Implement all `visitX` methods for each device
3. Pass it to `device.Accept(newVisitor)`

### Key files (quick glance)
- `IOTvisitor.go`
  - Visitor contract: `visitComputer`, `visitServer`, `visitPhone`, `visitPrinter`
- `IOTconnector.go`
  - Implements all visits and prints `Connecting ... connected`
- Device files
  - Each `Accept` calls the corresponding visit method

### Requirements
- Go 1.25+

That’s it. Simple, direct Visitor pattern with real-world flavor.


