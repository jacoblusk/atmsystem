package atmsystem

import (
	"errors"
)

type Bank struct {
	storage Storage
}

func (b *Bank) Deposit(id, amount int) error {
	a, err := b.storage.GetAccount(id)
	if err != nil {
		return err
	}

	err = b.storage.UpdateBalance(id, a.Balance+amount)
	return err
}

func (b *Bank) Withdrawl(id, amount int) error {
	a, err := b.storage.GetAccount(id)
	if err != nil {
		return err
	}

	if amount > a.Balance {
		return errors.New("Balance exceeded")
	}
	err = b.storage.UpdateBalance(id, a.Balance+amount)
	return err
}
