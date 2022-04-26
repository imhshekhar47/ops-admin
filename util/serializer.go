package util

import (
	"encoding/json"
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

func Serialize(message proto.Message) string {
	bytes, err := json.Marshal(message)
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

func Deserialize(jsonstr string, v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarhsal to proto.Message, message is %T", v)
	}

	return json.Unmarshal([]byte(jsonstr), p)
}
