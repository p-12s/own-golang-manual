package main

import "fmt"

type User struct { // recursion struct
	Name string
	parent *User
}

func main() {
	s := struct{}{}
	fmt.Printf("%T\n", s)

	i := 12
	fmt.Printf("%T\n", i)

	user := &User{}
	fmt.Printf("T user: %T\n", user)
	fmt.Println(user)

	user2 := User{}
	fmt.Printf("v user2: %v\n", user2)
	fmt.Println(user2)

	user3 := User {
		Name: "Ivan",
		parent: &user2,
	}
	fmt.Printf("v user3: %v\n", user3)
	fmt.Println(user3)
}
