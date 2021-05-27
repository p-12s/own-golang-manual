package main

import (
	"./calendar"
	"fmt"
	"log"
)

func main() {

	date := calendar.Date{}

	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err2 := date.SetMonth(5)
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := date.SetDay(11)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println(date.Day())
}
