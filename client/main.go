package main

import (
	"fmt"
	"github.com/jacoblusk/atmsystem"
	"time"
)

func main() {
	var amount int
	client, err := atmsystem.NewClient("localhost:1234", time.Second*60)
	if err != nil {
		fmt.Print(err)
		return
	}
	amount, err = client.Inquiry(100)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("amount: %d", amount)
}
