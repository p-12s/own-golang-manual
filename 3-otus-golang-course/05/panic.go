package main

import (
	"fmt"
	"os"
)

func main() {
	// panic - овать не нужно!
	// желательно только в случае, если ошибку нельзя обработать
	var gopath = os.Getenv("GOPATH_____")
	if gopath == "" {
		panic("no value for GOPATH")
	} else {
		fmt.Println("GOPATH: ", gopath)
	}
}
