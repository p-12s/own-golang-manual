package models

import (
	"github.com/jameycribbs/hare"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) SetID(i int) {
	u.Id = i
}

func (u User) GetID() int {
	return u.Id
}

func (u User) AfterFind(database *hare.Database) error {
	return nil
}
