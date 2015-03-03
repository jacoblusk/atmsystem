package atmsystem

//We use bson to serialize the object as leveldb is a very low level storage system and requires both key and value to be a byte array. So we take advantage of bson's ability to serialize go structures and use this for our value.

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID      int `bson:"id"`
	Balance int `bson:"balance"`
}

func (a *Account) MarshalBSON() ([]byte, error) {
	return bson.Marshal(a)
}

func (a *Account) UnmarshalBSON(b []byte) error {
	return bson.Unmarshal(b, a)
}
