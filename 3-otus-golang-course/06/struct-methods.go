package main

import "fmt"

type User struct {
	Id int64
	Name string
	Age int
	friends []int64
}

func (u User) IsOk() bool {
	for _, fid := range u.friends {
		if u.Id == fid {
			return true
		}
	}
	return false
}

func main() {
	var u = User{
		Id: 1,
		friends: []int64{1},
	}
	fmt.Println(u.IsOk())
}
