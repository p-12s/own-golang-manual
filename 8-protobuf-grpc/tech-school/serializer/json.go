package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) string {
	marshaller := protojson.MarshalOptions{
		Indent:          " ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	return marshaller.Format(message)
}
