package main

import (
	"github.com/own-golang-manual/0-golang-test-assignment/zhashkevych/coding-interview-challenge/generator"
	"fmt"
	"log"
)

func PrintTwoDimensionalArray(arr generator.Array) {
	for rowIndex, col := range arr {
		fmt.Printf("%d. %d\n", (rowIndex + 1), col)
	}
}

func main() {
	arr := generator.UniqueNumbers2Array {
		RowCount: 5,
		ColumnCount: 5,
		LeftRangeNumber: 1,
		RightRangeNumber: 25,
	}
	array, err := arr.Create()
	if err != nil {
		log.Fatal(err.Error())
	}
	PrintTwoDimensionalArray(array)
}
