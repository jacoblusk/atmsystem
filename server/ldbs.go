package main

import (
	"github.com/jacoblusk/atmsystem"
	"github.com/syndtr/goleveldb/leveldb"
	"encoding/binary"
)

type LDBStorage struct {
	DB *leveldb.DB
}

func(ldbs *LDBStorage) Open(filename string) error {
	var err error
	ldbs.DB, err = leveldb.OpenFile(filename, nil)
}

func(ldbs *LDBStorage) Close() error {
	return ldbs.DB.Close()
}

func (ldbs *LDBStorage) PutAccount(a atmsystem.Account) error {
	bson, err := a.MarshalBSON()
	if err != nil {
		return err
	}
	
	bs := [4]byte
	binary.LittleEndian.PutUint32(bs, a.ID)
	err = ldbs.DB.Put(bs, bson, nil)
}

func(ldbs *LDBStorage) UpdateBalance(id int, int balance) error {
	bs := [4]byte
	binary.LittleEndian.PutUint32(bs, id)
	a, err := ldbs.GetAccount(id)
	if err != nil {
		return err
	}

	a.Balance = balance; //update the balance
	var bson []byte
	bson, err = a.MarshalBSON() //reserialize
	if err != nil {
		return err
	}

	batch := new(leveldb.Batch)
	batch.Delete(bs)
	batch.Put(bs, bson)
	err = ldbs.DB.Write(batch, nil)
	return err
}

func (ldbs *LDBStorage) GetAccount(id int) error {
	bs := [4]byte
	binary.LittleEndian.PutUint32(bs, id)
	data, err := ldbs.DB.Get(bs, nil)
	if err != nil {
		return err
	}

	account = new(atmsystem.Account)
	err = account.UnmarshalBSON(data)
	return account, err
}
