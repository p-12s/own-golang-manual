package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1
	foo := func() {
		fmt.Println("Hello!")
	}
	foo()

	// 2
	func() {
		fmt.Println("Hello 2!")
	}() // "Hello!"

	// 3
	var foo2 func() = func() {
		fmt.Println("Hello 3!")
	}
	foo2() // Hello

	// example
	people := []string{"long-word", "middle", "small"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people)
}