proto:
	protoc --proto_path=./schema --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false schema/*.proto

clean:
	rm api/*.go

lint:
	golangci-lint run

test:
	go test -v -race ./...

run:
	go run cmd/main.go

run_race:
	go run -race cmd/*.go
