package generator

import (
	"fmt"
	"github.com/own-golang-manual/0-golang-test-assignment/zhashkevych/coding-interview-challenge"
	"testing"
	"github.com/stretchr/testify/require"
	generator "number_generator.go"
)

func Test_ValidateSettings(t *testing.T) {

	testTable := []struct {
		name         string
		rowCount, columnCount, leftRangeNumber, rightRangeNumber int
		wantErr      bool
	}{
		{
			name: "Ok",
			rowCount: 5,
			columnCount: 5,
			leftRangeNumber: 1,
			rightRangeNumber: 100,
			wantErr: false,
		},
	}
	//fmt.Println(testTable)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			arr := generator.UniqueNumbers2Array {
				RowCount: testCase.rowCount,
				ColumnCount: testCase.columnCount,
				LeftRangeNumber: testCase.leftRangeNumber,
				RightRangeNumber: testCase.rightRangeNumber,
			}

			require.Equal(t, Count(s, 'e'), 2, "counting 'e' in "+s)
			fmt.Println(arr)
		})
	}
	/*
	func (c *UniqueNumbers2Array) validateSettings() error {
		availableNumbersCount := c.RightRangeNumber - c.LeftRangeNumber
		findNumbersCount := c.RowCount * c.ColumnCount

		if availableNumbersCount <= 0 || availableNumbersCount < findNumbersCount {
			errorStr := fmt.Sprintf("Невозможно сгенерировать уникальные числа: нужно сгенерировать уникальных чисел: %d всего доступно чисел: %d",
				findNumbersCount, availableNumbersCount)
			return errors.New(errorStr)
		}
		return nil
	}

	arr := generator.UniqueNumbers2Array {
			RowCount: 5,
			ColumnCount: 5,
			LeftRangeNumber: 5,
			RightRangeNumber: 100,
		}
		array, err := arr.Create()


	*/
}
