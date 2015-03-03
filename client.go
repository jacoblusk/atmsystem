package atmsystem

import (
	"net"
	"net/rpc"
	"time"
)

//Here we wrap our client object with methods for interfacing with the RPC interface that is hosted by the server, this abstracts the need for the client to need to know about the specific method names that the server provides.

type Client struct {
	connection *rpc.Client
}

func NewClient(addr string, timeout time.Duration) (*Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	return &Client{connection: rpc.NewClient(conn)}, nil
}

func (c *Client) Deposit(id, amount int) (int, error) {
	var reply Transaction
	err := c.connection.Call("Bank.Deposit", &Transaction{ID: id, Amount: amount}, &reply)
	if err != nil {
		return 0, err
	}
	return reply.Amount, nil
}

func (c *Client) Withdraw(id, amount int) (int, error) {
	var reply Transaction
	err := c.connection.Call("Bank.Withdraw", &Transaction{ID: id, Amount: amount}, &reply)
	if err != nil {
		return 0, err
	}
	return reply.Amount, nil
}

func (c *Client) Inquiry(id int) (int, error) {
	var reply Transaction
	err := c.connection.Call("Bank.Inquiry", id, &reply)
	if err != nil {
		return 0, err
	}
	return reply.Amount, nil
}
