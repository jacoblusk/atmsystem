package atmsystem

import (
	_ "net/rpc"
)

type Server struct {
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) Start() {
}
