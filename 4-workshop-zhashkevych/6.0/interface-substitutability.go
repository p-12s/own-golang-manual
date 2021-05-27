package main

import (
	"errors"
	"fmt"
)

// =================

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

// =================

// storage должен уметь делать вставку, получение и удаление сотрудников
// как — это уже вопрос реализации
type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

// =================

type memoryStorage struct {
	data map[int]employee
}

// map init constructor
func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with this id dosn't exists")
	}
	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

// =================

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Println("Insert OK: ", e.id)
	return nil
}

func (s *dumbStorage) get(id int) (employee, error) {
	fmt.Println("Get OK: ", id)
	return employee{id: id}, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Println("Delete OK: ", id)
	return nil
}

// =================

func main() {
	// объявляем без первонач инициализации
	var s storage
	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	// инициал. первым типом
	s = newMemoryStorage()
	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	// инициал. вторым
	s = newDumbStorage()
	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	// сбрасываем
	s = nil
	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	// use examples
	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))
	fmt.Println(ms)

	spawnEmployees(ds)
	fmt.Println(ds)
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
	}
}
