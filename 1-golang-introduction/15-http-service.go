package main

import (
	"./datafile"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Guestbook struct {
	SignatureCount int
	Signature []string
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signature, err := datafile.GetStrings("data-signatures.txt")
	check(err)

	html, err := template.ParseFiles("view.html")
	check(err)



	guestbook := Guestbook {
		SignatureCount: len(signature),
		Signature: signature,
	}

	err = html.Execute(writer, guestbook)
	check(err)
}

// переход к добавлению новой записи
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)
}

// добавление новой записи
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("data-signatures.txt", options, os.FileMode(0600))
	_, err = fmt.Fprintln(file, signature)
	check(err)

	err = file.Close()
	check(err)

	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
