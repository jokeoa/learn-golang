package main

type Server struct {
	Brand string
	Model string
	Price float64
}

func (s *Server) Accept(visitor IOTvisitor) {
	visitor.visitServer(s)
}
