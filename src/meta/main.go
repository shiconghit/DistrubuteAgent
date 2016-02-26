package main

import proto "github.com/golang/protobuf/proto"

import (
	"./g"
	"./http"
	"flag"
	"fmt"
	"log"
	"os"
	"./mysql"
	"../common/common_proto"
	"reflect"
	"strings"
	"../common/message_pack"
)


func main() {
	//	TODO just for text
	//encode

	msg := &common_proto.Helloworld{
		Id: proto.Int32(101),
		Str:proto.String("hello"),
	}
	buffer, err := proto.Marshal(msg)
	if err != nil{
		log.Println("failed parser: %s\n", err)
	}

	massagename := getMessageName(msg)
	encodebuf := message_pack.EncodeMessagePack(massagename, buffer)

	anything :=  message_pack.DecodeMessagePack(encodebuf)
	//	obj, ok := anything.(common_proto.Helloworld)
	//	obj := (*common_proto.Helloworld)(anything)

	t := reflect.TypeOf(anything)

	log.Println(t)


	//	if obj, ok := anything.(string){
	//		log.Println(obj)
	//	}else{
	//		log.Println(ok)
	//	}

	os.Exit(0)

	//decode
	aa := common_proto.Helloworld{}
	rrstring := (reflect.TypeOf(&aa)).Elem().String()
	log.Println(rrstring)

	tt := proto.MessageType(strings.ToLower(rrstring))
	log.Println(tt)
	s := reflect.New(tt.Elem())

	msedecode := s.Interface().(proto.Message) //&common_proto.Helloworld{}
	err = proto.Unmarshal(buffer, msedecode)
	if err != nil{
		log.Println("failed parser: %s\n", err)
	}
	log.Println("decode: ", msedecode.String())

	os.Exit(0)
	//------------------------------------

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if err := g.ParseConfig(*cfg); err != nil {
		log.Println(err)
	}

	mysql.InitDb()

	go http.Start()

	select {}
}
