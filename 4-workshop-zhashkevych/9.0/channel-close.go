package main

import "fmt"

func main() {
	numbers := make(chan int)

	go generateNumbers(1000, numbers)

	for n := range numbers {
		fmt.Println(n)
	}
}

func generateNumbers(n int, res chan int) {
	for i := 0; i <= n; i++ {
		res <- i * 2
	}
	close(res) // if remove close - "fatal error: all goroutines are asleep - deadlock!"
}
