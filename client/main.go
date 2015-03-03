package main

import (
	"fmt"
	"github.com/jacoblusk/atmsystem"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	arglen := len(args)
	if arglen <= 4 {
		log.Fatalf("not enough arguments %d", arglen)
	}
	laddr := fmt.Sprintf("%s:%s", args[1], args[2])
	client, err := atmsystem.NewClient(laddr, time.Second*60)

	//Very ugly argument parsing, but it works!
	switch args[3] {
	case "inquiry":
		var id, balance int
		id, err = strconv.Atoi(args[4])
		if err != nil {
			log.Fatal("invalid id")
		}
		balance, err = client.Inquiry(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The current balance of account %d is $%d", id, balance)
		break
	case "deposit":
		var id, amount, balance int
		id, err = strconv.Atoi(args[4])
		if err != nil {
			log.Fatal("invalid id")
		}
		if arglen != 6 {
			log.Fatal("not enough arguments")
		}
		amount, err = strconv.Atoi(args[5])
		if err != nil {
			log.Fatal("invalid amount")
		}
		balance, err = client.Deposit(id, amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully deposited $%d to account %d, remaining balance $%d", id, amount, balance)
		break
	case "withdraw":
		var id, amount, balance int
		id, err = strconv.Atoi(args[4])
		if err != nil {
			log.Fatal("invalid id")
		}
		if arglen != 6 {
			log.Fatal("not enough arguments")
		}
		amount, err = strconv.Atoi(args[5])
		if err != nil {
			log.Fatal("invalid amount")
		}
		balance, err = client.Withdraw(id, amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully withdrew $%d from account %d, remaining balance $%d", id, amount, balance)
		break

	default:
		log.Fatalf("invalid command %s", args[2])
	}
}
