# Профилирование

Создание файлов:
```
go test -bench=UserSearch -benchmem -cpuprofile=cpu.out -memprofile=mem.out -x .
```
## 1 способ просмотра результата - консольный:
Сколько памяти было выделено:
```
go tool pprof -alloc_space mem.out
top
```
Команда top покажет результат:
```
Showing nodes accounting for 1406.62MB, 99.76% of 1409.98MB total
Dropped 28 nodes (cum <= 7.05MB)
flat  flat%   sum%        cum   cum%
1018.09MB 72.21% 72.21%  1018.09MB 72.21%  ../http-profiling.doUserSearch
388.54MB 27.56% 99.76%  1406.62MB 99.76%  ../http-profiling.BenchmarkUserSearch
0     0% 99.76%  1406.62MB 99.76%  testing.(*B).launch
0     0% 99.76%  1406.62MB 99.76%  testing.(*B).runN
```
больше всего памяти аллоцировал метод doUserSearch (72.21%), за ним BenchmarkUserSearch (27.56%).  
Взглянем на метод глубже
```
(pprof) list BenchmarkUserSearch

Total: 1.38GB
ROUTINE ======================== ../http-profiling.BenchmarkUserSearch in ../http-profiling/main_test.go
  388.54MB     1.37GB (flat, cum) 99.76% of Total
         .          .     17:}
         .          .     18:
         .          .     19:func BenchmarkUserSearch(b *testing.B) {
         .          .     20:
         .          .     21:	for i := 0; i < b.N; i++ {
  388.54MB   388.54MB     22:		req := new(userSearchRequest)
         .  1018.09MB     23:		_ = doUserSearch(req, userSearchParams)
         .          .     24:	}
         .          .     25:}
         .          .     26:
         .          .     27:func BenchmarkUserSearchWithPool(b *testing.B) {
         .          .     28:	for i := 0; i < b.N; i++ {
```
В сумме метод аллоцировал 1.37GB, на "оберточную работу" метод потратил 388.54MB  
Оставшиеся 1018.09MB - работа метода doUserSearch.  
Пройдем глубже:  
```
(pprof) list doUserSearch
Total: 1.38GB
ROUTINE ======================== ../http-profiling.doUserSearch in ../http-profiling/main.go
 1018.09MB  1018.09MB (flat, cum) 72.21% of Total
         .          .    143:	req.limit = args.GetUintOrZero("limit")
         .          .    144:	req.offset = args.GetUintOrZero("offset")
         .          .    145:
         .          .    146:	req.results = req.results[:0]
         .          .    147:	for i := 0; i < 5; i++ {
 1018.09MB  1018.09MB    148:		req.results = append(req.results, "result example")
         .          .    149:	}
         .          .    150:
         .          .    151:	sink = req
         .          .    152:
         .          .    153:	return responseStub
```
выйдем (quit).  
Посмотрим сколько раз менеджер памяти выделял память:
```
go tool pprof -alloc_objects mem.out
(pprof) top

Showing nodes accounting for 22254973, 100% of 22262353 total
Dropped 28 nodes (cum <= 111311)
      flat  flat%   sum%        cum   cum%
  18011129 80.90% 80.90%   18011129 80.90%  ../http-profiling.doUserSearch
   4243844 19.06%   100%   22254973   100%  ../http-profiling.BenchmarkUserSearch
         0     0%   100%   22254973   100%  testing.(*B).launch
         0     0%   100%   22254973   100%  testing.(*B).runN
         
(pprof) list doUserSearch

Total: 22262353
ROUTINE ======================== ../http-profiling.doUserSearch in ../http-profiling/main.go
  18011129   18011129 (flat, cum) 80.90% of Total
         .          .    143:	req.limit = args.GetUintOrZero("limit")
         .          .    144:	req.offset = args.GetUintOrZero("offset")
         .          .    145:
         .          .    146:	req.results = req.results[:0]
         .          .    147:	for i := 0; i < 5; i++ {
  18011129   18011129    148:		req.results = append(req.results, "result example")
         .          .    149:	}
         .          .    150:
         .          .    151:	sink = req
         .          .    152:
         .          .    153:	return responseStub
```
Разница только в метриках.  
Остался CPU-профайл:  
```
go tool pprof cpu.out

(pprof) top
Showing nodes accounting for 1950ms, 60.75% of 3210ms total
Dropped 39 nodes (cum <= 16.05ms)
Showing top 10 nodes out of 97
flat  flat%   sum%        cum   cum%
360ms 11.21% 11.21%     1890ms 58.88%  ../http-profiling.doUserSearch
300ms  9.35% 20.56%      300ms  9.35%  runtime.kevent
290ms  9.03% 29.60%      460ms 14.33%  github.com/valyala/fasthttp.peekArgStr (inline)
290ms  9.03% 38.63%      320ms  9.97%  runtime.heapBitsSetType
150ms  4.67% 43.30%      150ms  4.67%  runtime.pthread_cond_wait
120ms  3.74% 47.04%      120ms  3.74%  github.com/valyala/fasthttp.parseUintBuf
110ms  3.43% 50.47%      230ms  7.17%  github.com/valyala/fasthttp.ParseUint
110ms  3.43% 53.89%      110ms  3.43%  memeqbody
110ms  3.43% 57.32%      700ms 21.81%  runtime.mallocgc
110ms  3.43% 60.75%      110ms  3.43%  runtime.pthread_kill
```
Покажется, в каких функциях больше всего времени провела программа.  
(здесь часто показываются runtime-функции, GC).  
Также можно зайти внутрь функции и посмтореть детализацию:  
```
(pprof) list doUserSearch
Total: 3.21s
ROUTINE ======================== github.com/p-12s/own-golang-manual/3-otus-golang-course/17/http-profiling.doUserSearch in /Users/pavel/GitHub-repos/own-golang-manual/3-otus-golang-course/17/http-profiling/main.go
     360ms      1.89s (flat, cum) 58.88% of Total
         .          .    136:
         .          .    137:var responseStub = []byte(`{"organization": "a", "experience": 5}`)
         .          .    138:var sink *userSearchRequest
         .          .    139:
         .          .    140:func doUserSearch(req *userSearchRequest, args *fasthttp.Args) []byte {
      10ms      120ms    141:	req.name = args.Peek("name")
      20ms      130ms    142:	req.city = args.Peek("city")
         .      260ms    143:	req.limit = args.GetUintOrZero("limit")
         .      390ms    144:	req.offset = args.GetUintOrZero("offset")
         .          .    145:
         .          .    146:	req.results = req.results[:0]
      10ms       10ms    147:	for i := 0; i < 5; i++ {
     300ms      960ms    148:		req.results = append(req.results, "result example")
         .          .    149:	}
         .          .    150:
         .          .    151:	sink = req
         .          .    152:
      20ms       20ms    153:	return responseStub
     ...

quit
```
В примере ведущий показывает оптимизацию программы с помощью sync.Pool.
Но с оговоркой, что после такой оптимизации нужно проверить всю систему - улучшилась ли она?  
WRK в этом поможет.  

sync.Pool - список объектов из пула, Go будет пытаться вместо аллоцирования места в куче переиспользовать объекты из этого пула.  
За счет этого требования по-памяти уменьшаются.  
Однако нужно проверять результат оптимизации - иногда программа наоборот замедляется.  
Проблема пула в том, что GC может подчистить объекты из пула,  
и этим нарушив переиспользование объектов.  

Попробуем оценить результат профайлинга.  
Для объекивности запустим профайлер 10 раз,
для версии до оптимизации и после:
```
go test -bench=UserSearch$ -count=3 -benchmem -cpuprofile=cpu.old.out -memprofile=mem.old.out -x . | tee new.txt
go test -bench=UserSearchWithPool -count=3 -benchmem -cpuprofile=cpu.new.out -memprofile=mem.new.out -x . | tee new.txt
```
Мы перенаправили вывод в файлы:
```
goos: darwin
goarch: amd64
pkg: github.com/p-12s/own-golang-manual/3-otus-golang-course/17/http-profiling
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkUserSearchWithPool-16    	14046236	        83.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkUserSearchWithPool-16    	14004130	        77.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkUserSearchWithPool-16    	14674398	        79.38 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/p-12s/own-golang-manual/3-otus-golang-course/17/http-profiling	3.899s
```
Далее ставим **benchstat**.  
Команда (у меня не установился - настройки $GOROOT???):
```
go get golang.org/x/perf/cmd/benchstat
```
В моем примере названия функций разные, в боевых условиях предполагается, что они одинаковые - вы же запускаете бенчмарки над одними и теми же методами.  
Поэтоу я приведу названия типов к одному виду - "BenchmarkUserSearch-16".  
Результат оптимизации выведет команда:  
```
benchstat old.txt new.txt
```
delta - на сколько улучшилась/ухудшилась программа.  
p - статистический шум. чем ближе к "0" - тем лучше.  
У меня установить **benchstat** не вышло.  
   
Автор препарируемой программы добавил прямо в код возможность вызова профайлеров из строки браузера:
```
	switch string(ctx.Path()) {
	case "/userSearch":
		h.handleUserSearch(ctx)
	case "/startProfiling":
		h.startProfiling()
	case "/stopProfiling":
		h.stopProfiling()
	case "/stop":
		h.stop <- true
	default:
		ctx.Error("unknown resource accessed", fasthttp.StatusNotFound)
	}
```
предположу, что это используется там, где нет возможности подключиться к машине через консоль.  
Также он положил скрипт **run_wrk.bash** кторый автоматизирует запуск бенчмарков.  
Запуск программы выглядит так.  
Старая версия кода:
```
Прога:
go run main.go -memprofile=old_mem.out -cpuprofile-old_cpu.out

Нагрузка:
./run_wrk.bash -c100 -d10s -t4 http://127.0.0.1:8080/userSearch
```
Новая версия кода:
```
Прога:
go run main.go -memprofile=old_mem.out -cpuprofile-old_cpu.out -withPool=true

Нагрузка:
./run_wrk.bash -c100 -d10s -t4 http://127.0.0.1:8080/userSearch
```

## 2 способ результата - web view
```
go tool pprof -http=":8080" -alloc_space mem.new.out
go tool pprof -http=":8080" -alloc_objects mem.out
go tool pprof -http=":8080" cpu.out
```