package atmsystem

import (
	"net"
	"net/rpc"
)

type Server struct {
	bank     *Bank
	listener net.Listener
}

func NewServer(storage Storage) *Server {
	s := new(Server)
	s.bank = NewBank(storage)
	return s
}

func (s *Server) Start(laddr string) (err error) {
	rpc.Register(s.bank)
	s.listener, err = net.Listen("tcp", laddr)
	if err != nil {
		return nil
	}
	rpc.Accept(s.listener)
	return nil
}
