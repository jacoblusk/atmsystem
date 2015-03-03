package atmsystem

//Here we define the Storage interface which is used by the server's package main to implement server specific storage.
type Storage interface {
	PutAccount(account *Account) error
	GetAccount(id int) (*Account, error)
	Exists(id int) (bool, error)
}
