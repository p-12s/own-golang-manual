//go:generate ./command.sh

package main

import "fmt"

func main() {
	fmt.Println("run any unix command in go:generate")
}

//go:generate -command ls -l
//go:generate -command bye echo "Goodbye, world!"
//go:generate bye
//go:generate go run generate.go

// go generate
// go generate -v
// go generate -x
// go generate -n
// go generate -run echo -x
// go generate -run echo -n
// go generate -run bye
