package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof" // <- включаем pprof
)

type Resp struct {
	Status int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!\n\n")
	for i := 0; i < 100; i++ {
		resp := &Resp{200}
		res, _ := json.Marshal(resp)
		fmt.Fprintln(w, res)
	}
	fmt.Fprintln(w, "Hello, world!\n\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
