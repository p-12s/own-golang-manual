package main

import "fmt"

func main() {
	fmt.Println("Start")

	// создание канала и получение из него данных
	fmt.Println(<-createChan(5)) // 5

	fmt.Println("End")
}

func createChan(n int) chan int {
	ch := make(chan int)    // создаем канал
	go func(){
		ch <- n      // отправляем данные в канал
	}()             // запускаем горутину
	return ch   // возвращаем канал
}