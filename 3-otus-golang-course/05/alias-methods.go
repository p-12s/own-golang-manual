package main

import "fmt"

type MyInt int // alias

func (i MyInt) Print1() { // method for alias - "НЕ изменяющий"
	i++
	fmt.Println(i)
}

func (i *MyInt) Print2() { // method for alias - * значит "изменяющий"
	*i++
	fmt.Println(*i)
}

func main() {
	var i MyInt = 3
	i.Print1()
	i.Print1()

	i.Print2()
	i.Print2()
}
