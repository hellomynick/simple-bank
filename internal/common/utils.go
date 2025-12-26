package common

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func GetEventName(event proto.Message) string {
	return string(event.ProtoReflect().Descriptor().Name())
}

func MustParseInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(fmt.Errorf("can not parse int64 from string %s: %v", s, err))
	}

	return value
}
