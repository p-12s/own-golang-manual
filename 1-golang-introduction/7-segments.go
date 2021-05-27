package main

import "fmt"

func main() {
	// Сегмент (под капотом - массив)
	var notes = make([]string, 7)
	notes[0] = "Do"
	fmt.Println(notes[0])

	fruits := []string{
		"apple",
		"mango",
	}
	fruits = append(fruits, "orange")
	fmt.Println(fruits)
}
