package main

import (
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/01/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func writeToFile(filepath string, mess proto.Message) error {
	out, err := proto.Marshal(mess)
	if err != nil {
		log.Fatalln("can't serialize to byte", err)
	}

	if err := ioutil.WriteFile(filepath, out, 0644); err != nil {
		log.Fatalln("can't write to file", err)
		return err
	}
	fmt.Println("write ok!")
	return nil
}

func writeStringToFile(filepath string, mess string) error {
	if err := ioutil.WriteFile(filepath, []byte(mess), 0644); err != nil {
		log.Fatalln("can't write to file", err)
		return err
	}
	fmt.Println("write ok!")
	return nil
}

func readFromFile(filepath string, mess proto.Message) error {
	in, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("can't read file", err)
		return err
	}

	err = proto.Unmarshal(in, mess)
	if err != nil {
		log.Fatalln("can't unmarshal data", err)
		return err
	}
	return nil
}

func getDate() *pb.Date {
	return &pb.Date{
		Year:  2022,
		Month: 12,
		Day:   30,
	}
}

func protobufToJSON(message proto.Message) string {
	marshaller := protojson.MarshalOptions{
		Indent:          " ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	return marshaller.Format(message)
}

func main() {
	// записываем в бинарный файл протобаф
	date := getDate()
	writeToFile("8-protobuf-grpc/udemy-protocol-buffers-3/01/data.bin", date)

	// записываем в json-файл
	strView := protobufToJSON(date)
	fmt.Println(strView)
	writeStringToFile("8-protobuf-grpc/udemy-protocol-buffers-3/01/data.json", strView)

	// считываем
	sm2 := &pb.Date{}
	fmt.Println("before read:")
	fmt.Println(sm2)
	readFromFile("8-protobuf-grpc/udemy-protocol-buffers-3/01/data.bin", sm2)
	fmt.Println("after read:")
	fmt.Println(sm2)
}
