package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func someHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "h h h h!")
}

type countHandler struct {
	mu sync.Mutex // guard
	n int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is $d\n", h.n)
}

/*
+ простой http сервер
- в отдельном файле с возможностью graceful-shutdown
- с вынесением хендлеров
*/

func main() {
	http.HandleFunc("/www", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello 1")
	})
	http.Handle("/eee", new(countHandler))
	http.HandleFunc("/rrr", someHandler)
	http.Handle("/rro", http.NotFoundHandler())

	http.NewServeMux()

	server := http.Server{
		Addr: ":4343",

		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
