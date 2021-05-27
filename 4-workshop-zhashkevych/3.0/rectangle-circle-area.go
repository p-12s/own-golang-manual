package main

/*
Перепишите программу для подсчета площади круга, используя число Пи из библиотеки math.
Напишите функции для подсчета площади и периметра прямоугольника и треугольника.
*/

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printCircleArea(-2)
	printCircleArea(2)

	printRectangleInfo(0, 10)
	printRectangleInfo(-1, 10)
	printRectangleInfo(2, 3)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr2")
	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным\n")
	}

	return float32(radius) * float32(radius) * math.Pi, nil
}

func printRectangleInfo(a int, b int) {
	rectanglePerimeter, err := calculateRectanglePerimeter(a, b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rectangleArea, err := calculateRectangleArea(a, b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Стороны прямоугольника: %d и %d см.\n", a, b)
	fmt.Printf("Периметр прямоугольника: %d см. кв.\n", rectanglePerimeter)
	fmt.Printf("Площадь прямоугольника: %d см. кв.\n", rectangleArea)
}

func calculateRectanglePerimeter(a int, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Стороны прямоугольника должны быть положительными\n")
	}

	return (a + b) * 2, nil
}

func calculateRectangleArea(a int, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Стороны прямоугольника должны быть положительными\n")
	}

	return a * b, nil
}
