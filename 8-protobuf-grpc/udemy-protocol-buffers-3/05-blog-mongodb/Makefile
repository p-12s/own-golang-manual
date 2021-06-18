gen:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false proto/*.proto

clean:
	rm pb/*.go


s:
	go run server/server.go

c:
	go run client/client.go
