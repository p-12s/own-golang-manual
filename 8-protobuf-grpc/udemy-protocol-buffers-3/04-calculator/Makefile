gen:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false proto/*.proto

clean:
	rm pb/*.go


server:
	go run server/server.go

client:
	go run client/client.go
