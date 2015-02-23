package atmsystem

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
