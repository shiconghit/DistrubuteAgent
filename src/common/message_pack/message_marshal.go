package message_pack

import proto "github.com/golang/protobuf/proto"

import (
	"log"
	"hash/crc32"
	"bytes"
	"encoding/binary"
	"reflect"
	"errors"
)

func GetMessageName(msg interface{}) string{
	return (reflect.TypeOf(msg)).Elem().String()
}

func EncodeMessagePack(massagename string, protobufferdata []byte) []byte{
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

func DecodeMessagePack(buf []byte)(reobj interface{}, decode_error error){
	messagelen := binary.LittleEndian.Uint32(buf[0:4])
	log.Println("messagelen ", messagelen)
	log.Println("uint32(len(buf)) ", uint32(len(buf)))

	if(uint32(len(buf)) != messagelen){
		log.Println("length error")
		decode_error = errors.New("length error")
		return
	}

	chceksumget := binary.LittleEndian.Uint32(buf[len(buf)-4:])
	var checksum uint32 = crc32.ChecksumIEEE(buf[:len(buf)-4])
	if (chceksumget == checksum){
		var namelen uint32 = binary.LittleEndian.Uint32(buf[4:8])
		massagename := string(buf[8:8+namelen])
		log.Println("message name is" + massagename)

		messageType := proto.MessageType(massagename)
		if messageType == nil{
			log.Println("error proto.MessageType ")
			decode_error = errors.New("error proto.MessageType")
			reobj = nil
			return
		}
		log.Println(messageType)
		messageObj := reflect.New(messageType.Elem())

		msgdecode := messageObj.Interface().(proto.Message) //&common_proto.Helloworld{}
		err := proto.Unmarshal(buf[8+namelen: len(buf)-4], msgdecode)

		if err != nil{
			log.Println("failed parser: %s\n", err)
		}
		log.Println("decode: ", msgdecode.String())

		reobj = msgdecode
		return
	}else {
		log.Println("checksum error")
		reobj = nil
		return
	}
}