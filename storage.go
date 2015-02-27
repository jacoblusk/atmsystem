package atmsystem

type Storage interface {
	Open(filename string) error
	Close() error
	PutAccount(account Account) error
	GetAccount(id int) (Account, error)
	UpdateBalance(id, balance int) error
}
