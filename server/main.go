package main

import (
	"flag"
	"fmt"
	"github.com/jacoblusk/atmsystem"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const dbFilename = "bank_ldb_data"

func main() {
	var port int
	flag.IntVar(&port, "port", 1234, "port to start server on")
	ldbs := new(LDBStorage)
	err := ldbs.Open(dbFilename)
	defer ldbs.Close()
	if err != nil {
		log.Fatal(err)
	}

	//close the database file on sigterm/interrupt ^c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("interrupt detected, closing %s", dbFilename)
		ldbs.Close()
		os.Exit(0)
	}()

	var ok bool
	ok, err = ldbs.Exists(100)
	if err != nil {
		log.Fatal(err)
	}

	//keep the account persistent
	if !ok {
		account := new(atmsystem.Account)
		account.ID = 100
		account.Balance = 1000
		err = ldbs.PutAccount(account)
		if err != nil {
			ldbs.Close()
			log.Fatal(err)
		}
	}

	server := atmsystem.NewServer(ldbs)
	err = server.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Print(err)
	}
}
