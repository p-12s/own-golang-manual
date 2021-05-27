package main

import (
	"./storage"
	"fmt"
)

func spawnEmployees(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		s.insert(storage.Employee{id: i})
	}
}

func main() {
	ms := storage.newMemoryStorage()
	ds := storage.newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))
	fmt.Println(ms)

	spawnEmployees(ds)
	fmt.Println(ds)
}
