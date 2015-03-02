package atmsystem

import (
	"net"
	"net/rpc"
)

type Server struct {
	Bank     *Bank
	Listener net.Listener
}

func NewServer(storage Storage) *Server {
	s := new(Server)
	s.Bank = new(Bank)
	s.Bank.Storage = storage
	return s
}

func (s *Server) Start(laddr string) (err error) {
	rpc.Register(s.Bank)
	s.Listener, err = net.Listen("tcp", laddr)
	if err != nil {
		return nil
	}
	rpc.Accept(s.Listener)
	return nil
}
