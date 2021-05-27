package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id      int64   `json:"-"` // игнорировать в encode/json
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"user_age" db:"how_old"`
	friends []int64 `json:"friends"`
}

func main() {
	p1 := Person{
		Id:      1,
		Name:    "Ivan",
		Age:     32,
		friends: []int64{1},
	}
	fmt.Println(p1)

	js, _ := json.Marshal(p1)
	fmt.Printf("%s\n", string(js))
}
