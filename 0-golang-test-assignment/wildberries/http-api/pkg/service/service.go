package service

import (
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/repository"
)

type Authorization interface {

}

type Comment interface {
}

type Service struct {
	Authorization
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
