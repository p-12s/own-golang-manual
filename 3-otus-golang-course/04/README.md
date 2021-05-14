# Arrays and slices

## Array (use in matrix, for example)
```go
var arr [256]int        // фиксированная длинна
var arr [10][10]string  // может быть многомерным 
var arr [...]{1 ,2, 3}  // так пишут когда кол-во элементов в коде меняется
arr := [10]int{1,2,3,4,5} // остальное - нули

arr[3:5] // это уже слайс
```

## Slice
```go
var s []int   // неинициализированный слайс, nil
s := []int{}  // c помощью литерала слайса
s := make([]int, 3) // с помощью функции make, s == {0,0,0}
s := make([]int, 3, 10)

// base operations
s[i] = 1               // работает если i < len(s)
s[len(s) + 10] = 1     // случится panic
s = append(s, 1)       // добавляет 1 в конец слайса
s = append(s, 1, 2, 3) // добавляет 1, 2, 3 в конец слайса
s = append(s, s2...)   // добавляет содержимое слайса s2 в конец s
var s []int            // s == nil
s = append(s, 1)       // s == {1}  append умеет работать с nil слайсами
```

## Map (dictionary)
```go
var cache map[string]string     // не инициализированный словарь, nil
cache := map[string]string{}    // с помощью литерала, len(cache) == 0
cache := map[string]string{     // литерал с первоначальным значением
  "one":   "один",
  "two":   "два",
  "three": "три",
}
cache := make(map[string]string)  // тоже что и map[int]string{}
cache := make(map[string]string, 100)  // заранее выделить память 
                                       // на 100 ключей

// Map working
value := cache[key]     // получение значения, 
// если ключ не найден - Zero Value
value, ok := cache[key] // получить значение, и флаг того что ключ найден
_, ok := cache[key]     // проверить наличие ключа в словаре
cache[key] = value      // записать значение в инициализированный(!) словарь 
delete(cache, key)      // удалить ключ из словаря, работает всегда

// Map iteration (items randomized!)
for key, val := range cache {
...
}
for key, _ := range cache {  // eсли значение не нужно
...
}
for _, val := range cache {  // если ключ не нужен
...
}

// Get keys
var keys []string
for key, _ := range cache {
    keys = append(keys, key)
}

// Get values
values := make([]string, 0, len(cache)) // заранее выделим capacity
for _, val := range cache {
    values = append(values, val) // ускорение - append не будет выделять память
}

```