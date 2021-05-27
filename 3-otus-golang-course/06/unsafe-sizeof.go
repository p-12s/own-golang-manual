package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%d\n", unsafe.Sizeof(1)) // 8 на моей машине

	fmt.Printf("%d\n", unsafe.Sizeof("Asdfasdfdsf")) // 16 (длина + указатель)
	// строка в Go - структура из 2 полей: 1) длина строки uint64 2) указатель на массив где сама строка

	var x struct { // из-за смещений размер структуры в памяти может быть разным
		a byte   // 1
		b bool   // 1
		c uint64 // 8

	}
	fmt.Printf("%d\n", unsafe.Sizeof(x)) // 16 !

	fmt.Printf("%d\n", unsafe.Offsetof(x.a))
	fmt.Printf("%d\n", unsafe.Offsetof(x.c))
	fmt.Printf("%d\n", unsafe.Offsetof(x.b))
}
