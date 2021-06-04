package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

// brew install protoc-gen-go
// protoc -h
// go:generate protoc --go_out=. --csharp_out=. --cpp_out=. --java_out=. --js_out=. --kotlin_out=. --php_out=. --python_out=. --ruby_out=. Person.proto
//go:generate protoc --go_out=. Person.proto

func main() {
	p := new(Person)
	p.Name = "Anton"
	p.Mobile = append(p.Mobile, "8800553535")

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{}
	err = proto.Unmarshal(data, &p1)
	if err !=
		nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", p1)
}
