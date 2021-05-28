package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (d Dog) Say() {
	fmt.Println("ow-ow")
}

func main() {
	intB, _ := json.Marshal(1)
	fmt.Println("intB", intB)

	fltB, _ := json.Marshal(2.34)
	fmt.Println("fltB", fltB)

	strB, _ := json.Marshal("gopher")
	fmt.Println("strB", string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("slcD", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("mapD", string(mapB))

	dog := new(Dog)
	dog.Name = "Mursik"
	dog.Age = 21

	buff, err := json.Marshal(dog)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(buff)
	fmt.Println(string(buff))

	str := `{"Name":"Rax", "Age":32}`
	res := Dog{}
	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Start struct: %v %T\n", dog, dog)
	fmt.Printf("Unmarshal: string to struct: %v %T\n", res, res)

	//========
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v %T\n", animals, animals)
	for i, animal := range animals {
		fmt.Println(i, animal.Name, animal.Order)
	}
}
