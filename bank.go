package atmsystem

import (
	"errors"
)

type Bank struct {
	Storage Storage
}

type Transaction struct {
	ID, Amount int
}

func (b *Bank) Deposit(t, r *Transaction) error {
	a, err := b.Storage.GetAccount(t.ID)
	if err != nil {
		return err
	}

	newBalance := a.Balance + t.Amount

	r.ID = a.ID
	r.Amount = newBalance

	err = b.Storage.UpdateBalance(a.ID, newBalance)
	return err
}

func (b *Bank) Withdrawl(t, r *Transaction) error {
	a, err := b.Storage.GetAccount(t.ID)
	if err != nil {
		return err
	}

	if t.Amount > a.Balance {
		return errors.New("Balance exceeded")
	}

	newBalance := a.Balance - t.Amount
	r.ID = a.ID
	r.Amount = newBalance
	err = b.Storage.UpdateBalance(a.ID, newBalance)
	return err
}

func (b *Bank) Inquiry(id int, r *Transaction) error {
	a, err := b.Storage.GetAccount(id)
	if err != nil {
		return err
	}

	r.ID = a.ID
	r.Amount = a.Balance
	return err
}
