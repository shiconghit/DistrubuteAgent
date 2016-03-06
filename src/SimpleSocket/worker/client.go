package main

import (
	"time"
	"fmt"
	"net"
	"strconv"
)

func main() {
	c,err := net.Dial("unix", "/Volumes/HardWare/GOLearning/DistrubuteAgent/src/SimpleSocket/ipc/echo.sock")
	if err != nil {
		panic(err.Error())
	}
	for {
//		_,err := c.Write([]byte("hi\n"))
//		if err != nil {
//			println(err.Error())
//		}
//		time.Sleep(1e9)
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}
		data := buf[0:nr]
		fmt.Printf("Received: %v", string(data))

		x, err := c.Write([]byte("hi" + string(data) + "\n"))
		if err != nil {
			println(err.Error() + strconv.Itoa(x))
		}
		time.Sleep(1e9)
	}
}
