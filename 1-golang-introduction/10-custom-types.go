package main

import "fmt"

// Мой собственный тип данных на основе структуры - subscriber
// Для глобальной области видимости - всед.б с Большой
type Subscriber struct {
	Name string
	Rate float64
	Active bool
}

func printInfo (s Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Active: ", s.Active)
}

func defaultSubscribe (name string) Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 12.33
	s.Active = true
	return s
}

func applyDiscount (s *Subscriber) {
	s.Rate = 3.99
}

func main() {

	// Создание экземпляра кастомного типа
	subscriber1 := defaultSubscribe("Alex")
	printInfo(subscriber1)

	// изменение поля с пом. указателя
	applyDiscount(&subscriber1)
	printInfo(subscriber1)
}