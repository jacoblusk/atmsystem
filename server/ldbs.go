package main

import (
	"encoding/binary"
	"github.com/jacoblusk/atmsystem"
	"github.com/syndtr/goleveldb/leveldb"
)

type LDBStorage struct {
	DB *leveldb.DB
}

func (ldbs *LDBStorage) Open(filename string) error {
	var err error
	ldbs.DB, err = leveldb.OpenFile(filename, nil)
	return err
}

func (ldbs *LDBStorage) Close() error {
	return ldbs.DB.Close()
}

func (ldbs *LDBStorage) PutAccount(a *atmsystem.Account) error {
	bson, err := a.MarshalBSON()
	if err != nil {
		return err
	}

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(a.ID))
	err = ldbs.DB.Put(bs, bson, nil)
	return err
}

func (ldbs *LDBStorage) DeleteAccount(id int) error {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(id))
	err := ldbs.DB.Delete(bs, nil)
	return err
}

func (ldbs *LDBStorage) GetAccount(id int) (*atmsystem.Account, error) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(id))
	data, err := ldbs.DB.Get(bs, nil)
	if err != nil {
		return nil, err
	}

	account := new(atmsystem.Account)
	err = account.UnmarshalBSON(data)
	return account, err
}

func (ldbs *LDBStorage) Exists(id int) (bool, error) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(id))
	return ldbs.DB.Has(bs, nil)
}
