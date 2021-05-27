package main

import (
	"fmt"
)

func main() {

	// Длинная запись
	var notes [7]string
	notes[0] = "DO"
	fmt.Println(notes[0])

	// Запись через литералы
	notes2 := [7]string{
		"Do",
		"Re",
		"Mi",
		"Fa",
		"Sol",
		"La",
		"Si",
	}
	for i := 0; i < len(notes2); i++ {
		fmt.Println(notes2[i])
	}
	for index, value := range notes2 {
		fmt.Println(index, " ", value)
	}

}
