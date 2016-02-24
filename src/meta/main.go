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
)

func main() {
//	TODO just for text
	//encode
	tt := proto.MessageType("common_proto.helloworld")
	s := reflect.New(tt)
	log.Println(s.String())

	msg := &common_proto.Helloworld{
		Id: proto.Int32(101),
		Str:proto.String("hello"),
	}
	buffer, err := proto.Marshal(msg)
	if err != nil{
		log.Println("failed parser: %s\n", err)
	}
	log.Println(buffer)

	//decode
	msedecode := &common_proto.Helloworld{}
	err = proto.Unmarshal(buffer, msedecode)
	if err != nil{
		log.Println("failed parser: %s\n", err)
	}
	log.Println("decode: %s", msedecode.String())

	os.Exit(0)
//

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
