package main

type Employee struct {
	name, surname string
}

func FullName(e Employee) string {
	return e.name + " " + e.surname + "\n"
}

func (e Employee) FullName() string { // METHOD!!!
	return e.name + " " + e.surname + "\n"
}

func main() {
	print(Employee{"alexander", "davydov"}.FullName())
	print(FullName(Employee{"alexander", "davydov"}))
}
