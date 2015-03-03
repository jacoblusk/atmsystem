package atmsystem

import (
	"net"
	"net/rpc"
)

//The server contains an instance of bank, and the net.Listener interface. As of now, this server isn't stoppable; but could be by implementing a stoppable listener.
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
	//Here we register our bank instance as an rpc interface.
	rpc.Register(s.bank)
	s.listener, err = net.Listen("tcp", laddr)
	if err != nil {
		return nil
	}
	rpc.Accept(s.listener)
	return nil
}
