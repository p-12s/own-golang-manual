package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting ", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	/*sizes := make(chan int)
	go responseSize("https://golang.org/", sizes) // горутины
	go responseSize("https://golang.org/doc", sizes)

	fmt.Println(<-sizes)
	fmt.Println(<-sizes)*/

	// 1) способ дождаться завершения горутин
	// time.Sleep(3 * time.Second) // если main() завершится раньше чем отработают горутины - горутины убьются

	// 2) каналы - передача данных между горутинами
	// myChannel := make(chan float64)
	// myChannel <- 3.14 - добавление в канал
	// asss := <-myChannel - долучение из канала

	// перепишем через сегменты (цикл)
	pages := make(chan Page)
	urls := []string{
		"https://golang.org/",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println(page.URL, " ", page.Size)
	}
}
