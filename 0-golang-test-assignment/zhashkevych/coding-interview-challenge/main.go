package main

import (
	"github.com/own-golang-manual/0-golang-test-assignment/zhashkevych/coding-interview-challenge/generator"
	"fmt"
	"log"
)

func PrintTwoDimensionalArray(array generator.Array) error {
	for rowIndex, col := range array {
		fmt.Printf("%d. %d\n", (rowIndex + 1), col)
	}
	return nil
}

func main() {
	arr := generator.UniqueNumbers2Array {
		RowCount: 2,
		ColumnCount: 2,
		LeftRangeNumber: 1,
		RightRangeNumber: 4,
	}
	array, err := arr.Create()
	if err != nil {
		log.Fatal(err.Error())
	}
	PrintTwoDimensionalArray(array)
}
