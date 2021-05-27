package main

import "fmt"

func main() {
	var a int = 1
	var b *int = &a
	fmt.Printf("%d %d\n", a, *b)
	a++ // увеличит обе переменные
	fmt.Printf("%d %d\n", a, *b)
	*b++ // увеличит обе переменные
	fmt.Printf("%d %d\n", a, *b)

	c := 100
	b = &c // изменит c, оставив a нетронутой
	fmt.Printf("%d %d\n", a, *b)
}
