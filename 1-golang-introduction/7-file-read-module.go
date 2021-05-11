package main

import (
	"./datafile"
	"fmt"
	"log"
)

func main() {

	numbers, err := datafile.GetFloats("data-numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	count := float64(len(numbers))
	fmt.Println("Стреднее: ", sum / count)
}

