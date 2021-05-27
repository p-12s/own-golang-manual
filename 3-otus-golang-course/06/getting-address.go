package main

import "fmt"

func main() {
	var x struct {
		a int
		b string
		c [10]rune
	}
	bPtr := &x.b // получение адреса
	fmt.Println(bPtr)

	c3Ptr := &x.c[2] // получение адреса
	fmt.Println(c3Ptr)
}
