package message_pack

import proto "github.com/golang/protobuf/proto"

import (
//	"log"
//	"hash/crc32"
//	"bytes"
//	"encoding/binary"
//	"reflect"
//	"strings"
	"../common_proto"
	"testing"
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Name string
	Age  int
	Id   string
}

func (u User)sayHello() {
	fmt.Println("hello!")
}

func Test_case1(t *testing.T){
	u := &User{Name:"mckee", Age:20, Id:"user100"};

	t.Log(reflect.TypeOf(u))

	a := reflect.TypeOf(*u)
	t.Log(a) //main.User
	t.Log(a.Name()) //User
	t.Log(a.NumField()) //3
	t.Log(a.Kind()) //struct
}

func Test_case2(t *testing.T){
	msg := &common_proto.Helloworld{
		Id: proto.Int32(101),
		Str:proto.String("hello"),
	}
	t.Log(msg.Str, msg.Id)
	buffer, err := proto.Marshal(msg)
	if err != nil{
		t.Log("failed parser: %s\n", err)
	}

	massagename := GetMessageName(msg)
	encodebuf := EncodeMessagePack(massagename, buffer)

	obj, err :=  DecodeMessagePack(encodebuf)
	if err != nil{
		t.Log(err)
		return
	}
	t.Log(reflect.TypeOf(obj))
	k := obj.(proto.Message).(*common_proto.Helloworld)

	t.Log(k.GetId(), k.GetStr(), k.GetOpt())

//	if(reflect.TypeOf(obj) == reflect.TypeOf(&common_proto.Helloworld{})){
//		reflect.Ptr(obj)
//	}

	//	re := reflect.ValueOf(obj)
	//	//	obj, ok := anything.(common_proto.Helloworld)
	//	//	obj := (*common_proto.Helloworld)(anything)
	//	msgdecode := (re.Interface()).(proto.Message)
	//	t.Log( msgdecode )
//	obj := reflect.TypeOf(*anything)

//	t.Log(obj.Id(), obj.Str())
}
//
func Test_case3(t *testing.T){
	msg := &common_proto.Helloworld{
		Id: proto.Int32(101),
		Str:proto.String("hello"),
	}
	t.Log(msg.Str, msg.Id)
	buffer, err := proto.Marshal(msg)
	if err != nil{
		t.Log("failed parser: %s\n", err)
	}

	massagename := GetMessageName(msg)
	encodebuf := EncodeMessagePack(massagename, buffer)

	obj, err :=  DecodeMessagePack(encodebuf)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(reflect.TypeOf(obj))
	k := obj.(proto.Message).(*common_proto.Helloworld)

	t.Log(k.GetId(), k.GetStr(), k.GetOpt())

	//	if(reflect.TypeOf(obj) == reflect.TypeOf(&common_proto.Helloworld{})){
	//		reflect.Ptr(obj)
	//	}

	//	re := reflect.ValueOf(obj)
	//	//	obj, ok := anything.(common_proto.Helloworld)
	//	//	obj := (*common_proto.Helloworld)(anything)
	//	msgdecode := (re.Interface()).(proto.Message)
	//	t.Log( msgdecode )
	//	obj := reflect.TypeOf(*anything)

	//	t.Log(obj.Id(), obj.Str())
}


var chhh = make(chan common_proto.Helloworld, 128)
func start(){
	for{
		msg := common_proto.Helloworld{
			Id: proto.Int32(101),
			Str:proto.String("hello"),
		}
		chhh <- msg
		time.Sleep(time.Microsecond * 100)
	}

}

func Test_case4(t *testing.T) {
	go start()
	for{
		t.Log(<-chhh)
	}

}