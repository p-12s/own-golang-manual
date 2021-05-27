package main

import (
	"fmt"
	"time"
)

/*
Поищите дополнительную информацию
про буферизированные и однонаправленные каналы.
*/

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for { // бесконечный цикл, пишущий в канал через интервал
			time.Sleep(time.Millisecond * 500)
			message1 <- "Times up 0.5 sec."
		}
	}()

	go func() {
		for { // бесконечный цикл, пишущий в канал через интервал
			time.Sleep(time.Second * 2)
			message2 <- "Times up 2 sec."
		}
	}()

	/*for {
		fmt.Println(<-message1)
		fmt.Println(<-message2) // чтение будет блокироваться на 2 канале, вывод будет по-очереди
	}
	*/

	// для ситуаций с блокировками при чтении, когда в 1 канал уже пришли данные но мы не можем их прочитать есть select -
	// для неблокирующего чтения из нескольких каналов

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}
	}
	// теперь ОК
}
