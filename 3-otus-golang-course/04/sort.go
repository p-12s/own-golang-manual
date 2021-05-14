package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

func main() {
	s := []int{0, 10, 5, 9, 6, 8, 3, 2, 1}
	fmt.Println("Before: ", s)
	sort.Ints(s)
	fmt.Println("After: ", s)

	s2 := []string{"a", "z", "zzza", "bbb", "hello", "cruel", "world"}
	fmt.Println("Before: ", s2)
	sort.Strings(s2)
	fmt.Println("After: ", s2)

	// а что если нужно сортировать свои типы ?
	s3 := []User{
		{"vasya", 19},
		{"Djo", 49},
		{"petya", 18},
		{"Ivan", 30},
		{"Ivan", 34},
	}
	fmt.Println("Before: ", s3)
	sort.Slice(s3, func(i, j int) bool {
		return s3[i].Age < s3[j].Age
	})
	fmt.Println("After: ", s3)

}
