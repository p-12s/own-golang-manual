gen:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. proto/*.proto

gen2:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false proto/*.proto


clean:
	rm pb/*.go
