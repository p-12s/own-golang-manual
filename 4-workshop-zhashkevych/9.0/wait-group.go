package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
Изучите самостоятельно структуры
для блокировок и синхронизации горутин
sync.Mutex и sync.WaitGroup из стандартной библиотеки.
Поищите примеры их применения на практике.
 */

func main() {
	urls := []string{
		"https://google.com",
		"https://youtube.com",
		"https://vk.com",
		"https://ok.ru",
		"https://github.com",
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

func doHTTP(url string) {
	t := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get <%s> %s\n", url, err.Error())
	}

	defer resp.Body.Close()
	fmt.Printf("<%s> - Status [%d] - Latency %d ms\n",
		url, resp.StatusCode, time.Since(t).Milliseconds())
}
