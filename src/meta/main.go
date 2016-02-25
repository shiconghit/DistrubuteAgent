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
	"hash/crc32"
	"bytes"
	"encoding/binary"
)

func encodeMessagePack(massagename string, protobufferdata []byte) []byte{
	//简单的通信协议
	buf := new(bytes.Buffer)

	// namelen 打包进去解析的名称的长度
	var namelen int = len(massagename)

	// len 数据包的长度
	var messagelen int = 4 + 4 + namelen + len(protobufferdata) + 4
	log.Println("messagelen ", messagelen)

	//依次填充包
	err := binary.Write(buf,binary.LittleEndian, uint32(messagelen))
	if err != nil {
		log.Println(err)
	}
	log.Println(buf.Bytes())

	binary.Write(buf,binary.LittleEndian, uint32(namelen))
	log.Println(buf.Bytes())

	binary.Write(buf,binary.LittleEndian, []byte(massagename))
	log.Println(buf.Bytes())

	//protobufdata
	binary.Write(buf,binary.LittleEndian, protobufferdata)

	// checksum 		unint32
	var checksum uint32 = crc32.ChecksumIEEE(buf.Bytes())
	binary.Write(buf,binary.LittleEndian, checksum)

	log.Println("checksum ", checksum)

	return buf.Bytes()
}

func decodeMessagePack(buf []byte) interface{}{
	messagelen := binary.LittleEndian.Uint32(buf[0:4])
	log.Println("messagelen ", messagelen)
	log.Println("uint32(len(buf)) ", uint32(len(buf)))

	if(uint32(len(buf)) != messagelen){
		log.Println("length error")
		return nil
	}

	chceksumget := binary.LittleEndian.Uint32(buf[len(buf)-4:])
	var checksum uint32 = crc32.ChecksumIEEE(buf[:len(buf)-4])
	if (chceksumget == checksum){
		var namelen uint32 = binary.LittleEndian.Uint32(buf[4:8])
		massagename := string(buf[8:8+namelen])
		log.Println("message name is" + massagename)

		messageType := proto.MessageType(strings.ToLower(massagename))
		log.Println(messageType)
		messageObj := reflect.New(messageType.Elem())

		msgdecode := messageObj.Interface().(proto.Message) //&common_proto.Helloworld{}
		err := proto.Unmarshal(buf[8+namelen: len(buf)-4], msgdecode)

		if err != nil{
			log.Println("failed parser: %s\n", err)
		}
		log.Println("decode: ", msgdecode.String())

		return msgdecode
	}else {
		log.Println("checksum error")
		return nil
	}
}

func getMessageName(msg interface{}) string{
	return (reflect.TypeOf(msg)).Elem().String()
}

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
	encodebuf := encodeMessagePack(massagename, buffer)

	decodeMessagePack(encodebuf)
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
