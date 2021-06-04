package student

type Discipline = string

//go:generate /Library/go/go1.16.4/bin/bin/easyjson -all student.go
type Student struct {
	FirstName  string
	SecondName string
	Age        int
	Marks      map[Discipline]int
}
