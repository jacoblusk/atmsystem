package server

type Server struct {
}

func NewServer() *Server {
	Bank bank
}

func (s *Server) Start() {
	rpc.Register()
}
