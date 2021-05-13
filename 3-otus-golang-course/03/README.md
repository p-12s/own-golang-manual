## Work with UTF-8
Packages unicode и unicode/utf8
```go
// получить первую руну из строки и ее размер в байтах
DecodeRuneInString(s string) (r rune, size int)

// получить длинну строки в рунах
RuneCountInString(s string) (n int)

// проверить валидность строки
ValidString(s string) bool
```

## Standart library for string-work (strings)
```go
// проверка наличия подстроки
Contains(s, substr string) bool

// строка начинается с ?
HasPrefix(s, prefix string) bool

// склейка строк
Join(a []string, sep string) string

// разбиение по разделителю
Split(s, sep string) []string
```
For optimizing memory allocation:
```go
import "strings"
var b strings.Builder
for i := 33; i >= 1; i­­ {
    b.WriteString("Код")
    b.WriteRune('ъ')
}
result := b.String()
```
## String iteration by byte
```go
for i := 0; i < len(s); i++ {
    b := s[i]
    // i строго последоваельно
    // b имеет тип byte, uint8
}
```

## String iteration by rune
```go
for i, r := range s {
    // i может перепрыгивать значения 1,2,4,6,9...
    // r ­ имеет тип rune, int32
}
```

## Constants
```go
const PI = 3             // принимает подходящий тип
const pi float32 = 3.14  // строгий тип
const (
    TheA = 1
    TheB = 2
)
const (
    X = iota   // 0
    Y          // 1
    Z          // 2
)
```