package repository

type Authorization interface {
}

type Comment interface {
}

type Repository struct {
	Authorization
	Comment
}

func NewRepository() *Repository {
	return &Repository{}
}
