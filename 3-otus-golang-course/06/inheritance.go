package main

import "fmt"

type A struct {
	Field int
}

type B struct { // аналог наследования
	A
	Field int
}

func (a *A) GetField() int {
	return a.Field
}

func main() {
	a := &B{}

	a.A.Field = 200
	a.Field = 100

	fmt.Printf("%v %v %v", a, a.Field, a.GetField())
}