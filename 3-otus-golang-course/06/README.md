## Structures

```go
struct {} // пустая структура, не занимает памяти

type User struct { // структура с именованными полями 
    Id int64
    Name string
    Age int
    friends []int64 // приватный элемент
}

// init
u1 := User {} // Zero Value для типа User
u2 := &User {} // Тоже, но указатель
u3 := User { 1, "Vasya", 23 } // По номерам полей
u4 := User {
    Id: 1,
    Name: "Vasya",
    friends: []int64{1, 2, 3},
}
```

## Pointers - exists, address arithmetic - NOT
```go
x := 1 // Тип int
xPtr := &x // Тип *int
var p *int // Тип *int, значение nil

// getting address
var x struct {
    a int
    b string
    c [10]rune
}
bPtr := &x.b // получение адреса
fmt.Println(bPtr)

c3Ptr := &x.c[2] // получение адреса
fmt.Println(c3Ptr)
```
## Dereferencing pointers
```go
a := "qwe" // Тип string
aPtr := &a // Тип *string

b := *aPtr // Тип string, значение "qwe"
var n *int // nil
nv := *n // panic

p := struct{x, y int }{1, 3} // структура
pPtr= &p // указатель
fmt.Println(pPtr.x)
fmt.Println(pPtr.y)


// Методы объявленные над типом получают копию объекта, поэтому не могут его изменять!
func (u User) HappyBirthday() {
    u.Age++ // это изменение будет потеряно
}
// Методы объявленные над указателем на тип - могут
func (u *User) HappyBirthday() {
    u.Age++ // OK
}
```

## Inheritance analog
```go
type A struct {
    Field int
}

type B struct { // аналог наследования
    A
    Field int
}
```

## Build-in structures 
```go
type LinkStorage struct {
    sync.Mutex // только тип!
    storage map[string]string // тип и имя
}

storage := LinkStorage{}
storage.Mutex.Lock() // имя типа используется
storage.Mutex.Unlock() // как имя элемента структуры
```

## Tags
```go
// соглашение о структуре
`key:"value" key1:"value1,value11"`

type User struct {
    Id int64 `json:"-"` // игнорировать в encode/json
    Name string `json:"name"`
    Age int `json:"user_age" db:"how_old:`
    friends []int64
}

// Получить информацию о тэгах можно через reflect
u := User{}
ut := reflect.TypeOf(u)
ageField := ut.FieldByName("Age")
jsonSettings := ageField.Get("json") // "user_age"
```