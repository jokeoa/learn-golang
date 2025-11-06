package main

// Server represents a server device.
type Server struct {
	Brand string
	Model string
	Price float64
}

// Accept lets a visitor operate on the Server.
func (s *Server) Accept(visitor IOTvisitor) {
	visitor.visitServer(s)
}
