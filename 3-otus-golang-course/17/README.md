# Профилирование

## Best practices
Преаллокация slices, map всегда, когда известна вместительность
```go
// bad!!
const cap = 32000
a := make([]int, 0)
for i := 0; i < cap; i++ {
	a = append(a, i)
}
// OK
const cap = 32000
a := make([]int, 0, cap)
for i := 0; i < cap; i++ {
    a = append(a, i)
}
```
Переиспользование объектов
```go
type A struct {
	I int
}
func (a *A) Reset() {
	*a = A{}
}
// bad
for i := 0; i < 1000; i++ {
	a := &A{}
	json.Unmarshal([]byte("{\"i\":32}"), a)
}
// OK
a := &A{}
for i := 0; i < 1000; i++ {
	a.Reset()
    json.Unmarshal([]byte("{\"i\":32}"), a)
}
```
Конвертация числа в строку
```go
// bad
fmt.Sprintf("%d", 42) // reflection
// OK
strconv.FormatInt(42)
```
Поиск подстроки в строке
```go
var re *regexp.Regexp
// Excellent
strings.Contains("Hello world!", "world")
// Slow (хотя бы компилировать регулярку заранее, не в цикле)
res, _ := regexp.MatchString("world", "Hello world!")
return res
```
StringBuilder
```go
func Slow() string {
	a := ""
	a += "h"
	a += "i"
	return a
}
func Fast() string {
	builder := strings.Builder{}
	builder.WriteString("h")
	builder.WriteString("i")
	return builder.String()
}
func VeryFast() string { // на стеке
    return "h"+"i"
}
```
sync.Pool (шарим данные между горутинами) - пример возможно не точен на посл. версии Go
```go
var pool = sync.Pool {
    New: func() interface{} {
        return &strings.Builder{}
    },
}
func Slow() string {
    builder := strings.Builder{}
    builder.WriteString("Hello")
    return builder.String()
}
func Fast() string {
    builder := pool.Get().(*strings.Builder)
    defer pool.Put(builder)  // Try to comment it out!
    builder.Reset()
    builder.WriteString("Hello")
    res := builder.String()
    return res
}
```
Держать в heap (куче) много структур с указателями - дорого
```go
var arr1 = map[string]int{} // дорого
var arr2 = map[int][]byte{} // дорого
var arr3 = map[int]int{} // дешевле
```
Переиспользование HTTP-соединений
```go
var client = http.Client{
    Timeout:   3 * time.Second,
    Transport: &http.Transport{},
}
func Fast() ([]byte, error) {
    resp, err := client.Get("https://localhost")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return body, nil
}
```
## Бенчмаркинг
Пишем тест:
```go
import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getComDomains("data.dat") // тестируемый метод
	}
}
```
Запускаем тест-бенчмарк:
```
go test -bench=. -cpu=8 -benchmem -cpuprofile=cpu.out -memprofile=mem.out . 
```
WRK (если это http-сервис)
```
wrk -c100 -d10s -t50 http://127.0.0.1:8080/
c - connections
d - durations
t - threads
```
## Профайлинг
0) Самый минимальный профайлинг - проверить как выделяется память:
```
go build -gcflags '-m' ./main.go
```
1) Пример более полного. 
Генерируем файлы профайлинга (-bench=. здесь точка - регулярка, принимающая название файлов):
```
go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out -x .
```
Запускаем утилиту анализа для CPU.
Если искать по-назвнаию метода/объекта:
```
go tool pprof bench.test cpu.out

help
list Slow
list Fast
disasm Fast (дизассемблер)
quit
```
Запускаем утилиту анализа для Memory:
```
go tool pprof bench.test mem.out

help
list Slow
list Fast
quit
```
Профилируем HTTP-сервис.
запускаем сервис:  
```
go run main.go
```
запускаем нагрузку (чтобы профайлинг не был пустым):
```
wrk -c100 -d2000s -t50 http://127.0.0.1:8080/
```
запускаем профайлер (5 сек):
```
go tool pprof -http=":9090" -seconds=5 http://localhost:8080/debug/pprof/profile
```
в результате в браузере откроется граф выполнения.  
Отключаем нагрузку.  
Как читать граф:  
**https://github.com/google/pprof/blob/master/doc/README.md#interpreting-the-callgraph**

P.S.  
Если бы я не использовал http в своем приложнии, можно было бы запустить http отдельно:
```go
import (
	"net/http"
	_ "net/http/pprof"
)

func main () {
    http.ListenAndServe(":8081", nil)
}
```
P.P.S.
Не получилось открыть в браузере 127.0.0.1:8081/debug/pprof/ в остальном - профайлинг сработал.  

## Трейсинг
Выглядит так, но я не пробовал его (если приспичит - надо будет пересмотреть):
```
wrk -c100 -d2000s -t50 http://127.0.0.1:8080/
wget http://localhost:8080/debug/pprof/trace?seconds=5 -o trace.out
go tool trace trace.out
```

## Linux profiler
http://www.brendangregg.com/perf.html
