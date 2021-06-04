package main

import (
	"fmt"
)

var VersionString int = 0

func main() {
	fmt.Println("Version:", VersionString)
}

// go run -ldflags '-X main.VersionString=1.0' main.go
