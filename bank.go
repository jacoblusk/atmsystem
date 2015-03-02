package atmsystem

import (
	"errors"
	"sync"
)

type Bank struct {
	storage Storage
	mutex   *sync.Mutex
}

func NewBank(s Storage) *Bank {
	b := new(Bank)
	b.storage = s
	b.mutex = new(sync.Mutex)
	return b
}

type Transaction struct {
	ID, Amount int
}

func (b *Bank) Deposit(t, reply *Transaction) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	account, err := b.storage.GetAccount(t.ID)
	if err != nil {
		return err
	}

	if t.Amount < 0 {
		return errors.New("amount less than 0")
	}

	account.Balance += t.Amount
	reply.ID = account.ID
	reply.Amount = account.Balance

	err = b.storage.PutAccount(account)
	return err
}

func (b *Bank) Withdraw(t, reply *Transaction) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	account, err := b.storage.GetAccount(t.ID)
	if err != nil {
		return err
	}

	if t.Amount > account.Balance {
		return errors.New("balance exceeded")
	}

	account.Balance -= t.Amount
	reply.ID = account.ID
	reply.Amount = account.Balance

	err = b.storage.PutAccount(account)
	return err
}

func (b *Bank) Inquiry(id int, r *Transaction) error {
	a, err := b.storage.GetAccount(id)
	if err != nil {
		return err
	}

	r.ID = a.ID
	r.Amount = a.Balance
	return err
}
