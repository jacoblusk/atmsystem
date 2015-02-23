package atmsystem

type Storage interface {
	Open(filename string) error
	Close() error
	PutAccount(account atmsystem.Account) error
	GetAccount(id int) (atmsystem.Account, error)
	UpdateBalance(id int, balance int) error
}
