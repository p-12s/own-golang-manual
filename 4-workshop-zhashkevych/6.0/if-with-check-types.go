package main

import "fmt"

func printType(value interface{}) {
	if _, ok := value.(int); ok {
		fmt.Println(value, " - int")
	} else if _, ok := value.(string); ok {
		fmt.Println(value, " - string")
	} else {
		fmt.Println(value, " - other type")
	}
}

func printType2(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("тип аргумента int")
	case string:
		fmt.Println("тип аргумента string")
	default:
		fmt.Println("тип аргумента не int и не string")
	}
}

func main() {
	printType(3)
	printType("hello")
	printType(true)

	printType2("hello")
}
