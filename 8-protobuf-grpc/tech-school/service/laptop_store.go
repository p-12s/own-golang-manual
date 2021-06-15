package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/pb"
	"sync"
)

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

var ErrAlreadyExists = errors.New("record already exists")

func (i *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if i.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copier
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	i.data[other.Id] = other
	return nil
}

func (i *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	laptop := i.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
