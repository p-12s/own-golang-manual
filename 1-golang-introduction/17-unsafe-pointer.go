package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const (
	sizeUint8   = unsafe.Sizeof(uint8(0))
	sizeUint16  = unsafe.Sizeof(uint16(0))
	sizeUint32  = unsafe.Sizeof(uint32(0))
	sizeUint64  = unsafe.Sizeof(uint64(0))
	sizeUint    = unsafe.Sizeof(uint(0))
	sizeUintptr = unsafe.Sizeof(uintptr(0))

	sizeInt8  = unsafe.Sizeof(int8(0))
	sizeInt16 = unsafe.Sizeof(int16(0))
	sizeInt32 = unsafe.Sizeof(int32(0))
	sizeInt64 = unsafe.Sizeof(int64(0))
	sizeInt   = unsafe.Sizeof(int(0))

	sizeFloat32 = unsafe.Sizeof(float32(0))
	sizeFloat64 = unsafe.Sizeof(float64(0))

	sizeComplex64  = unsafe.Sizeof(complex64(0))
	sizeComplex128 = unsafe.Sizeof(complex128(0))

	sizeString = unsafe.Sizeof(string(""))
	sizeByte   = unsafe.Sizeof(byte(0))   // type byte = uint8
	sizeRune   = unsafe.Sizeof(rune('a')) // type rune = int32
)

func main() {
	// пример как посмотерть хранение в памяти BigEndian и LittleEndian
	v := uint16(1)
	big := make([]byte, 2)
	little := make([]byte, 2)

	binary.BigEndian.PutUint16(big, v)
	binary.LittleEndian.PutUint16(little, v)

	fmt.Println(big, little)
	// [0 1] [1 0]

	// просмотр кол-ва занимаемой памяти типами
	fmt.Println(sizeUint8, sizeUint16, sizeUint32, sizeUint64, sizeUint, sizeUintptr)
	// 1 2 4 8 8 8
	fmt.Println(sizeInt8, sizeInt16, sizeInt32, sizeInt64, sizeInt)
	// 1 2 4 8 8
	fmt.Println(sizeFloat32, sizeFloat64)
	// 4 8
	fmt.Println(sizeComplex64, sizeComplex128)
	// 8 16
	fmt.Println(sizeString, sizeByte, sizeRune) // type byte = uint8, type rune = int32
	// 16 1 4

	// ОПАСНОСТИ РАБОТЫ С unsafe.Pointer

	// Раньше при конвертации 2-х байтного числа в "более"-байтное в оставшиеся разряды записывались не 0
	// (похоже не стирались предыдущие значения из ячеек)
	// теперь все записывается 0-ми6 пофиксили
	aa := uint16(65535)
	bb := *(*[8]byte)(unsafe.Pointer(&aa))
	fmt.Println(bb)
	// 255 255 0 0 0 0 0 0

	// go vet отлавливает ошибки работы с модулем unsafe, но в сложных случаях не поможет
	a := 5
	e := unsafe.Pointer(&a)

	fmt.Println(e)
	// когда это нужно использ? - когда нужна супер оптимизацция
	// поискать в домашках отуса unsafe
	// поискать разборы задач с unsafe

}
