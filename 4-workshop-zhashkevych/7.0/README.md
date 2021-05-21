# Work with external modules

1) create main.go
2) run 
```
go mod init "github.com/p-12s/some-repo"
```   
3) import external mod:
```
 go get github.com/zhashkevych/scheduler
```
4) use mod:
```go
package main

import (
	"context"
	"fmt"
	"github.com/zhashkevych/scheduler"
	"time"
)

func main() {
	s := scheduler.NewScheduler()
	s.Add(context.Background(), func(ctx context.Context){
		fmt.Printf("Time is: %s\n", time.Now())
	}, time.Second * 1)
	time.Sleep(time.Minute)
}
```
5) to download all external modules on new computer:
```go
go mod download
```