package main

import "fmt"

type User struct {
	Name string // публичное поле
	password string // приватное
}


func main() {

	var storage map[string]string = make(map[string]string)
	// укоротим:
	var storage2 = make(map[string]string)

	fmt.Println(storage)
	fmt.Println(storage2)


	var i int = 10
	// укоротим:
	i2 := 10

	fmt.Println(i)
	fmt.Println(i2)
	
	/*
	Целые: int , uint , int8 , uint8 , ...
	Алиасы к целым: byte = uint8 , rune = int32
	С плавающей точкой: float32 , float64
	Комплексные: complex64 , complex128
	Строки: string
	Указатели: uintptr , *int , *string , ...
	*/
}
