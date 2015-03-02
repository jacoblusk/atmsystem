atmsystem
=========

Remote Invocation Implementation of A Simple ATM System

Submitted Files
---------------

account.go 		- structure for account  
bank.go    		- rpc interface  
client.go  		- structure and methods for client  
server.go  		- structure and methods for server  
storage.go 		- interface for storage  

server/main.go 	- executable for server  
server/ldbs.go 	- contains implementation for the storage interface  

client/main.go 	- executable for client

Building
--------

To install Go follow the guide [here](https://golang.org/doc/install)

One can obtain the source using the `go get` command to download and install  
`go get github.com/jacoblusk/atmsystem`

Or seperately one can download the source and use solong as their $GOPATH is
set one can use `go get` to install the dependinces and run using `go build`   
`go build github.com/jacoblusk/atmsystem/server`   
`go build github.com/jacoblusk/atmsystem/client`   
