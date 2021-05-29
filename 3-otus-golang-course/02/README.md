# Testing

UPD: гораздо лучшие примеры тестирования в этом проекте (параллельные, есть моки):
https://github.com/p-12s/todo-list-rest-api

## Method
```
package foobar

func Count(s string, r rune) int {
	var cnt int
	for _, l := range s {
		if l == r {
			cnt += 1
		}
	}
	return cnt
}
```

## Test
```
package foobar

import "testing"

func TestCount(t *testing.T) {
	s := "qwerasdfe"
	e := 2
	if c := Count(s, 'e'); c != e {
		t.Fatalf("bad count for %s: got %d expected %d", s, c , e)
	}
}
```

## Test running
```
cd 02
go test -cover (go tool cover -help)
```

### Package Testing
```
t.Fail()   // отметить тест как сломаный, но продолжит выполнение
t.FailNow()  // отметить тест как сломаный и прекратить текущий тест
t.Logf(formar string, ...interface{})  // вывести сообщение с отладкой
t.Errorf(formar string, ...interface{})  // t.Logf + t.Fail
t.Fatalf(formar string, ...interface{})  // t.Logf + t.FailNow
t.SkipNow()  // пропустить тест
```