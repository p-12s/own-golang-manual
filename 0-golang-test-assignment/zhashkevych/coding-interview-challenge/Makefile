.PHONY:
.SILENT:
.DEFAULT_GOAL := run

run:
	env GO111MODULE=off go run ./main.go

test:
	env GO111MODULE=off go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	env GO111MODULE=off go tool cover -func=cover.out

lint:
	/usr/local/Cellar/golangci-lint/1.40.1/bin/golangci-lint run
