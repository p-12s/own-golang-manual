package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func newHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, world!\n\n")
}

func main() {
	// GIN
	/*router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.Run(":80")*/

	// simple handler
	/*http.HandleFunc("/lo", newHandler)
	http.ListenAndServe(":80", nil)*/

	// more complecity
	http.HandleFunc("/lo2", newHandler)
	http.NewServeMux()

	server := http.Server{
		Addr:           ":81",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
