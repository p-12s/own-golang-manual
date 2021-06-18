# Evans

1. Добавляем в запсук сервера регистрацию reflection-сервиса:
```go
s := grpc.NewServer()
pb.RegisterCalculatorServiceServer(s, &server{})

// register reflection service
reflection.Register(s)
```
2. Устанавливаем утилиту Evans (репо https://github.com/ktr0731/evans):
```
brew tap ktr0731/evans
brew install evans
```
3. Запускаем сервер и evans (на порту сервера):
```
go run server.go
evans -p 50051 -r
```
4. Используем команды для просмотра доступных методов:
![Evans](https://github.com/p-12s/own-golang-manual/blob/master/8-protobuf-grpc/udemy-protocol-buffers-3/04-calculator/evans.png?raw=true)

5. Тестируем сервер RPC без клиента:
```
show service
service CalculatorService
call Sum (далее вводим данные и получаем результат)
```


