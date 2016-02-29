package main

import proto "github.com/golang/protobuf/proto"

import (
//	"log"
//	"hash/crc32"
//	"bytes"
//	"encoding/binary"
//	"reflect"
//	"strings"
	"./common_proto"
	"fmt"
	"time"
	"encoding/json"
)

var chhh = make(chan common_proto.Helloworld, 128)
func start(){
	count := 0
	for{
		count++
		str := fmt.Sprintf("hello %d", count)
		msg := common_proto.Helloworld{
			Id: proto.Int32(int32(count)),
			Str:proto.String("hello" + str),
		}
		chhh <- msg
		time.Sleep(time.Second)
	}

}

func main() {
	msg := common_proto.Helloworld{
		Id: proto.Int32(int32(100)),
		Str:proto.String("hello" + "dddd"),
	}
	by, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json marshal error")
	}
	fmt.Println(string(by))

	var xx common_proto.Helloworld;
	json.Unmarshal(by, &xx)

	fmt.Println(xx.GetId(), xx.GetStr())
}
