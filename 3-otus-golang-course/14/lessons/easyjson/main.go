package main

import (
	"encoding/json"
	"fmt"

	"2_8_codegen/easyjson/student"
)

func main() {

	s := student.Student{
		FirstName:  "Otus",
		SecondName: "Otusov",
		Age:        25,
		Marks: map[student.Discipline]int{
			"Golang":     5,
			"JavaScript": 3,
		},
	}
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}



