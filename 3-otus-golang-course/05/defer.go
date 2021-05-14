package main

import "fmt"

func c() (i int, j int) {
	defer func() { i *= 10 } ()
	return 10, 20
}

func main() {
	a, b := c()
	fmt.Println(a, "", b)
}
