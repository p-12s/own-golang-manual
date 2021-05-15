package main

import (
	"fmt"
)

type Stack struct {
	numbers []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.numbers) == 0
}

func (s *Stack) Push(i int) {
	s.numbers = append(s.numbers, i)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty ")
		return 0
	} else {
		index := len(s.numbers) - 1
		element := s.numbers[index]
		fmt.Println("Removed ", element)
		s.numbers = s.numbers[:index]
		return element
	}
}

func main() {
	s := Stack{}
	fmt.Println(s.IsEmpty())

	s.Push(10)
	fmt.Println(s)

	s.Push(20)
	fmt.Println(s)

	s.Push(30)
	fmt.Println(s)

	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
