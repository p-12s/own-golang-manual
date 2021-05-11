package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// Types
	var num2, num3 float64 = 2.31, 42.33223
	num4 := 42
	fmt.Println(num2, num3, num4)

	// Math
	fmt.Println(math.Floor(2.15))
	fmt.Println(strings.Title("Заголовок"))
	fmt.Println('A') // тип данных руны - символ в юникоде

	// Reflect
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.43))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("This is str"))
}
