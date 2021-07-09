FROM golang:latest
RUN mkdir app
ADD . ./app
WORKDIR ./app
RUN go mod tidy

ENV TableName="users"
ENV DataDir="./data"

ENTRYPOINT go run cmd/userService/main.go

EXPOSE 8080
