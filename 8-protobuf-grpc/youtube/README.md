# Сервер - клиент на gRPC

## 
Создаем директорию и proto-файл  
убеждаемся что плагины установлены

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc

Запускаем генерацию
protoc --go-grpc_out=. *.proto
protoc --go_out=. *.proto
protoc --go_out=. --go-grpc_out=. *.proto

/Library/go/go1.16.5/bin/bin

THIS SHIT DOSN'T WORK

Check this shit:
https://golang-blog.blogspot.com/2021/03/grpc-golang-basic-tutorial.html
 
возможно неполный файл
protoc --go_out=. --go-grpc_out=. proto/*.proto



