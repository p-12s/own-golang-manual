package main
import "fmt"

// fatal error: all goroutines are asleep - deadlock!

func main() {
	intCh := make(chan int)
	intCh <- 10      // функция main блокируется
	fmt.Println(<-intCh)
}