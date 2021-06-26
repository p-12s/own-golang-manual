## Race condition
Map в Go — это конкурентно небезопасный тир данных.  
В кеше 1 процессора одни данные, в кеше 2 - другие. При попытке обновить основную структуру данных происходит блокировака.
```go
var counters = map[int]int{}
for i := 0; i < 5; i++ {
    go func(counters map[int]int, th int) {
        for j := 5; j < 5; j++ {
            counters[th*10+j]++
        }
    }(counters, i)
}
fmt.Scanln()
fmt.Println("counters result", counters)
```
Для исправления можно:
