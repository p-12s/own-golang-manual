package storage

import (
	"errors"
	"fmt"
)

type Employee struct {
	id int
	name string
	age int
	salary int
}

// =================

// storage должен уметь делать вставку, получение и удаление сотрудников
// как — это уже вопрос реализации
type Storage interface {
	insert (e Employee) error
	get (id int) (Employee, error)
	delete (id int) error
}

// =================

type memoryStorage struct {
	data map[int]Employee
}

// map init constructor
func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *memoryStorage) insert(e Employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("employee with this id dosn't exists")
	}
	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

// =================

type dumbStorage struct {}

func newDumbStorage() * dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e Employee) error {
	fmt.Println("Insert OK: ", e.id)
	return nil
}

func (s *dumbStorage) get(id int) (Employee, error) {
	fmt.Println("Get OK: ", id)
	return Employee{id: id}, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Println("Delete OK: ", id)
	return nil
}

// =================

