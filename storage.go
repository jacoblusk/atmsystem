package atmsystem

type Storage interface {
	PutAccount(account *Account) error
	GetAccount(id int) (*Account, error)
	Exists(id int) (bool, error)
}
