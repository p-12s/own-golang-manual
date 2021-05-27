package main

import "fmt"

func main() {
	// канал только для отправки данных
	// var inCh chan<- int

	// канал только для получения данных
	// var outCh <-chan int

	intCh := make(chan int, 2)

	go factorial2(5, intCh)

	fmt.Println(<-intCh)
	fmt.Println("The End")
}

func factorial2(n int, ch chan<- int) { // канал, доступный только для отправки данных
	// внутри функции factorial мы можем только отправлять данные в канал, но не получать их
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}
