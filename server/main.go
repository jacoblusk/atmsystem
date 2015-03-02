package main

import (
	"flag"
	"fmt"
	"github.com/jacoblusk/atmsystem"
	"log"
)

const dbFilename = "bank.ldb"

func main() {
	var port int
	flag.IntVar(&port, "port", 1234, "port to start server on")
	ldbs := new(LDBStorage)
	err := ldbs.Open(dbFilename)
	defer ldbs.Close()
	if err != nil {
		log.Fatal(err)
	}

	account := new(atmsystem.Account)
	account.ID = 100
	account.Balance = 1000

	err = ldbs.PutAccount(account)
	if err != nil {
		ldbs.Close()
		log.Fatal(err)
	}

	server := atmsystem.NewServer(ldbs)
	err = server.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Print(err)
	}
}
