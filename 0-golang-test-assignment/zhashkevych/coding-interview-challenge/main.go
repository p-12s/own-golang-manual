package main

import (
	"./generator"
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
		RowCount: 5,
		ColumnCount: 5,
		LeftRangeNumber: 5,
		RightRangeNumber: 100,
	}
	array, err := arr.Create()
	if err != nil {
		log.Fatal(err.Error())
	}
	PrintTwoDimensionalArray(array)
}
