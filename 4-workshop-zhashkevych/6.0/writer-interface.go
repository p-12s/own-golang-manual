package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const message = "I use Go interfaces"

type httpWriter struct {
	client *http.Client
	addr   string
}

func newHttpWriter(addr string) *httpWriter {
	return &httpWriter{
		client: http.DefaultClient,
		addr:   addr,
	}
}

func (w httpWriter) Write(p []byte) (n int, err error) {
	_, err = w.client.Post(w.addr, "text/plain", bytes.NewBuffer(p))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	file, err := os.OpenFile("message.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	buffer := &bytes.Buffer{}
	httpWriter := newHttpWriter("http://localhost:1111/")

	writeMessageEncoded(file, message)
	writeMessageEncoded(buffer, message)
	writeMessageEncoded(httpWriter, message)

	fmt.Println("Buffer: ", buffer.String())
}

func writeMessageEncoded(writer io.Writer, message string) {
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	writer.Write([]byte(encodedMessage))
}
