package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UniqueNumbers2Array struct {
	RowCount int
	ColumnCount int
	LeftRangeNumber int
	RightRangeNumber int

	shuffledRange []int
	resultArr [][]int
}

// Валидация значений в конфиге
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

// Создание перемешанного массива
func (c *UniqueNumbers2Array) makeShuffledRange() {
	c.shuffledRange = make([]int, c.RightRangeNumber - c.LeftRangeNumber + 1)
	for i := range c.shuffledRange {
		c.shuffledRange[i] = c.LeftRangeNumber + i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c.shuffledRange), func(i, j int) {
		c.shuffledRange[i], c.shuffledRange[j] = c.shuffledRange[j], c.shuffledRange[i]
	})
}

type Array [][]int

// Заполнение двумерного массива из перемешанного одномерного массива
func (c *UniqueNumbers2Array) fillArray() {
	c.resultArr = make([][]int, c.RowCount)
	index := 0

	for row := 0; row < c.RowCount; row++ {
		c.resultArr[row] = make([]int, c.ColumnCount)

		for col := 0; col < c.ColumnCount; col++ {
			c.resultArr[row][col] = c.shuffledRange[index]
			index += 1
		}
	}
}

// Непосредственно создание массива
func (c *UniqueNumbers2Array) Create() (Array, error) {
	if err := c.validateSettings(); err != nil {
		return nil, err
	}

	c.makeShuffledRange()
	c.fillArray()

	return c.resultArr, nil
}