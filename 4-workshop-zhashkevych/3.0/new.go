package main

import "fmt"

/*
Найдите материал про функцию new() в Go,
и какое она имеет отношение к указателям

// https://rtfm.co.ua/golang-ukazateli-podrobnyj-razbor/
*/

func main() {
	a := 1
	b := new(int)
	fmt.Printf("a: %d %v | b: %v %v \n", a, &a, b, *b)
	// a: 1 0xc0000ae008 | b: 0xc0000ae010 0

	b = &a
	fmt.Printf("a: %d %v | b: %v %v \n", a, &a, b, *b)
	// a: 1 0xc0000ae008 | b: 0xc0000ae008 1
}
