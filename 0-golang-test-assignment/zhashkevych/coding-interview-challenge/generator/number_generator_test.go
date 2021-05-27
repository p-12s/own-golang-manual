package generator

import (
	"testing"
)

func Test_ValidateSettings(t *testing.T) {

	testTable := []struct {
		name         string
		rowCount, columnCount, leftRangeNumber, rightRangeNumber int
		wantErr      bool
	}{
		{
			name: "Ok if find elements < range",
			rowCount: 5,
			columnCount: 5,
			leftRangeNumber: 1,
			rightRangeNumber: 100,
			wantErr: false,
		},
		{
			name: "Ok if find elements = range",
			rowCount: 5,
			columnCount: 5,
			leftRangeNumber: 1,
			rightRangeNumber: 25,
			wantErr: false,
		},
		{
			name: "Not ok if find elements < range",
			rowCount: 5,
			columnCount: 5,
			leftRangeNumber: 1,
			rightRangeNumber: 24,
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			arr := UniqueNumbers2Array{
				RowCount: testCase.rowCount,
				ColumnCount: testCase.columnCount,
				LeftRangeNumber: testCase.leftRangeNumber,
				RightRangeNumber: testCase.rightRangeNumber,
			}

			err := arr.validateSettings()
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}

func Test_MakeShuffledRange(t *testing.T) {
	testTable := []struct {
		name         string
		rowCount, columnCount, leftRangeNumber, rightRangeNumber int
	}{
		{
			name:             "Creates an array 5x5 of unique shuffled numbers",
			rowCount:         5,
			columnCount:      5,
			leftRangeNumber:  1,
			rightRangeNumber: 25,
		},
		{
			name:             "Creates an array 10x10 of unique shuffled numbers",
			rowCount:         10,
			columnCount:      10,
			leftRangeNumber:  1,
			rightRangeNumber: 100,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			// Prepare
			arr := UniqueNumbers2Array{
				RowCount: testCase.rowCount,
				ColumnCount: testCase.columnCount,
				LeftRangeNumber: testCase.leftRangeNumber,
				RightRangeNumber: testCase.rightRangeNumber,
			}
			arr.validateSettings()
			arr.makeShuffledRange()

			// Checks
			checkNumbersUnique(t, &arr)
			checkNumbersShuffle(t, &arr)
		})
	}
}

func Test_FillArray(t *testing.T) {
	testTable := []struct {
		name         string
		rowCount, columnCount, leftRangeNumber, rightRangeNumber int
		wantErr      bool
	}{
		{
			name:             "5x5 ok",
			rowCount:         5,
			columnCount:      5,
			leftRangeNumber:  1,
			rightRangeNumber: 25,
		},
		{
			name:             "10x5 ok",
			rowCount:         10,
			columnCount:      5,
			leftRangeNumber:  1,
			rightRangeNumber: 50,
		},
		{
			name:             "10x10 ok",
			rowCount:         10,
			columnCount:      10,
			leftRangeNumber:  1,
			rightRangeNumber: 100,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Prepare
			arr := UniqueNumbers2Array{
				RowCount: testCase.rowCount,
				ColumnCount: testCase.columnCount,
				LeftRangeNumber: testCase.leftRangeNumber,
				RightRangeNumber: testCase.rightRangeNumber,
			}
			arr.validateSettings()
			arr.makeShuffledRange()
			arr.fillArray()

			// Check
			checkingArrayMatchesSource(t, &arr)
		})
	}
}

func Test_Create(t *testing.T) {
	testTable := []struct {
		name         string
		rowCount, columnCount, leftRangeNumber, rightRangeNumber int
		wantErr      bool
	}{
		{
			name:             "5x5 ok",
			rowCount:         5,
			columnCount:      5,
			leftRangeNumber:  1,
			rightRangeNumber: 25,
			wantErr:			false,
		},
		{
			name:             "10x5 ok",
			rowCount:         10,
			columnCount:      5,
			leftRangeNumber:  1,
			rightRangeNumber: 50,
			wantErr:			false,
		},
		{
			name:             "10x10 ok",
			rowCount:         10,
			columnCount:      10,
			leftRangeNumber:  1,
			rightRangeNumber: 100,
			wantErr:			false,
		},
		{
			name:             "10x10!=99 not ok",
			rowCount:         10,
			columnCount:      10,
			leftRangeNumber:  1,
			rightRangeNumber: 99,
			wantErr:			true,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Prepare
			arr := UniqueNumbers2Array{
				RowCount: testCase.rowCount,
				ColumnCount: testCase.columnCount,
				LeftRangeNumber: testCase.leftRangeNumber,
				RightRangeNumber: testCase.rightRangeNumber,
			}

			_, err := arr.Create()
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}

			if err == nil {
				checkingArrayMatchesSource(t, &arr)
			}
		})
	}
}

// ПРОВЕРИМ УНИКАЛЬНОСТЬ
func checkNumbersUnique(t *testing.T, arr *UniqueNumbers2Array) {
	// вспомогательный массив для хранения уникальных чисел
	uniqueNumbersMap := make(map[int]struct{}, (arr.RightRangeNumber - arr.LeftRangeNumber + 1))

	// поочередно берем элементы и проверяем, не существует ли уже такого в списке уникальных? если нет - добавляем его
	for i := 0; i < len(arr.shuffledRange); i++ {
		_, ok := uniqueNumbersMap[arr.shuffledRange[i]]
		if ok {
			t.Fatal("Items in array not unique:", i, arr.shuffledRange)
		}
		uniqueNumbersMap[arr.shuffledRange[i]] = struct{}{}
	}

	// для компрометации уникальности (проверка самих тестов) задваиваем первые 2 значения и ловим ошибку
	notUniqueNumbersMap := make(map[int]struct{}, arr.RightRangeNumber - arr.LeftRangeNumber + 1)

	item0 := arr.shuffledRange[0]
	arr.shuffledRange[0] = 1
	item1 := arr.shuffledRange[2]
	arr.shuffledRange[1] = 1

	findError := false
	for i := 0; i < len(arr.shuffledRange); i++ {
		_, ok := notUniqueNumbersMap[arr.shuffledRange[i]]
		if ok {
			findError = true
		}
		notUniqueNumbersMap[arr.shuffledRange[i]] = struct{}{}
	}
	if !findError {
		t.Fatal("Compromised not unique array is not throw error: ", arr.shuffledRange)
	}
	arr.shuffledRange[0] = item0
	arr.shuffledRange[1] = item1
}

// ПРОВЕРИМ ПЕРЕМЕШАННОСТЬ
func checkNumbersShuffle(t *testing.T, arr *UniqueNumbers2Array) {
	numbersInTheirOrdinalPlacesCount := 0 // кол-во чисел на своих изначально-правильных порядковых местах: 1 на 1, 2 на 2 и т.п.
	for i := 0; i < len(arr.shuffledRange); i++ {
		if i == arr.shuffledRange[i] {
			numbersInTheirOrdinalPlacesCount += 1
		}
	}
	// для простоты критерия считаем, что если более половины чисел остались на своих местах - ошибка
	// (для небольшого массива вероятность ошибки возрастает)
	if numbersInTheirOrdinalPlacesCount > len(arr.shuffledRange)/2 {
		t.Fatal("The numbers in the array are not jumbled: ", arr.shuffledRange)
	}
}

// ПРОВЕРИМ СООТВЕТСТВИЕ ЭЛЕМЕНТОВ МАССИВА ЭЛЕМЕНТАМ ПЕРЕМЕШАННОГО ИСХОДНИКА
func checkingArrayMatchesSource(t *testing.T, arr *UniqueNumbers2Array) {
	index := 0
	for i := 0; i < len(arr.resultArr); i++ {
		for j := 0; j < len(arr.resultArr[i]); j++ {
			if arr.resultArr[i][j] != arr.shuffledRange[index] {
				t.Fatal("Array elements do not match the sources:", arr.shuffledRange, arr.resultArr)
			}
			index++
		}
	}
}