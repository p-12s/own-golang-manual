# Concurrency

## 1. Gorutines
Channel can be closed:
```go
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
```
## 2. Channels
Non-buferized (without capacity):
```go
func main() {
	t := time.Now() // время старта

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	result1 := make(chan int)
	result2 := make(chan int)

	// очень сложные вычисления
	go calculateSomething(1000, result1)
	go calculateSomething(2000, result2)

	fmt.Println(<-result1)
	fmt.Println(<-result2)

	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}
```
Buferized (with capacity):
```go
func main() {
	intCh := make(chan int, 3)

	intCh <- 10
	intCh <- 3
	intCh <- 24
	// intCh <- 15  // блокировка - функция main ждет, когда освободится место в канале

	fmt.Println("capacity: ", cap(intCh))     // емкость элементов в канале
	fmt.Println("length: ", len(intCh))     // количество элементов в канале

	fmt.Println(<-intCh)     // 10
	fmt.Println(<-intCh)     // 3
	fmt.Println(<-intCh)     //24
}
```

## 3. Transfer of data streams
```go
package main

import "fmt"

func main(){
    intCh := make(chan int) 

    go factorial(7, intCh)

    for num := range intCh{
        fmt.Println(num)
    }
}

func factorial(n int, ch chan int){
    defer close(ch)
    result := 1
    for i := 1; i <= n; i++{
        result *= i
        ch <- result
    }
}
```
## 4. Mutex
```go
package main

import (
    "fmt"
    "sync"
)

var counter int = 0 // общий ресурс

func main() {
    ch := make(chan bool) // канал
    var mutex sync.Mutex // определяем мьютекс
    
    for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

    for i := 1; i < 5; i++ {
        <-ch
    }

    fmt.Println("The End")
}

func work (number int, ch chan bool, mutex *sync.Mutex) {
    mutex.Lock() // блокируем доступ к переменной counter
    counter = 0
    for k := 1; k <= 5; k++ {
        counter++
        fmt.Println("Goroutine", number, "-", counter)
    }
    mutex.Unlock() // деблокируем доступ
    ch <-true
}
```
## 5. WaitGroup  
```go
func main() {
	urls := []string{
		"https://google.com",
		"https://youtube.com",
	}

	var wg sync.WaitGroup // счетчик

	for _, url := range urls {
		wg.Add(1) // увеличиваем значение
		
		go func(url string) {
			doHTTP(url)
			
			wg.Done() // декрементит счетчик
		}(url)
	}

	wg.Wait() // блокирует, пока счетчик не равен 0
}
```

