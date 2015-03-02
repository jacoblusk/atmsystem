package atmsystem

import (
	"net"
	"net/rpc"
	"time"
)

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

func (c *Client) Withdrawl(id, amount int) (int, error) {
	var reply Transaction
	err := c.connection.Call("Bank.Withdrawl", &Transaction{ID: id, Amount: amount}, &reply)
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
