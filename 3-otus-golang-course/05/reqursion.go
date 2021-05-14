package main

import "fmt"

// n! = n×(n-1)! where n >0
func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	} else {
		return 1 // 1! == 1
	}
}

func main() {
	fmt.Println(getFactorial(10))
}