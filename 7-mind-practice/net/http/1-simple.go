package main

import (
	"fmt"
	"net/http" // - как для http-серверов, так и для киентов
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello simple")
	})
	http.ListenAndServe(":3435", nil)
}
