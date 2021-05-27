package main

import "fmt"

func main() {
	intCh := make(chan int) // небуферизированный канал - без указания емкости

	go factorial(45, intCh)
	fmt.Println(<-intCh)
	fmt.Println("The end")
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("%d! = %d\n", n, result)
	ch <- result
}
