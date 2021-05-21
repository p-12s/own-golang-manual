package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("Time: %s; Request: %s\n", time.Now().String(), body)
	})
	http.ListenAndServe(":1111", nil)
}
