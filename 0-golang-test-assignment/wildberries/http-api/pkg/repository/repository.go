package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api"
)

type Authorization interface {
	CreateUser(user common.User) (int, error)
}

type Comment interface {
}

type Repository struct {
	Authorization
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
