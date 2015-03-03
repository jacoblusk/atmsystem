package atmsystem

import (
	"errors"
	"sync"
)

//Here we use a mutex for Withdraw and Deposit as there is a possible race condition between checking if the account exists and depositing. This assignment doesn't have us deleting user acctions, but for scalability reasons it is used.
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

//This is the object that is passed as both a request and a reply, it could be abstracted further to have a seperate reply object for each RPC method, but that wasn't needed for this assignment.
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

	//Verify we aren't depositing negative money
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

	//Verify that we aren't overdrafting
	if t.Amount > account.Balance {
		return errors.New("balance exceeded")
	}

	if t.Amount < 0 {
		return errors.New("amount less than 0")
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
