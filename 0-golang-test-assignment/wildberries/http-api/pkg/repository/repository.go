package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Comment interface {
}

type Repository struct {
	Authorization
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
