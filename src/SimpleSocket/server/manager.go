package main

import (
	"net"
	"fmt"
//	"path/filepath"
)

func echoServer(c net.Conn) {
	for {
//		buf := make([]byte, 512)
//		nr, err := c.Read(buf)
//		if err != nil {
//			return
//		}
//		addr := filepath.Base(c.LocalAddr().String())
//		extension := filepath.Ext(addr)
//		proc := addr[0 : len(addr)-len(extension)]
//		fmt.Printf("%s %s %s", addr, extension, proc)

//		data := buf[0:nr]
//		fmt.Printf("Received: %v", string(data))
		_, err := c.Write([]byte("shicong"))
		if err != nil {
			panic("Write: " + err.Error())
		}

		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}
		data := buf[0:nr]
		fmt.Printf("Received: %v", string(data))
	}
}

func main() {
	l, err := net.Listen("unix", "/Volumes/HardWare/GOLearning/DistrubuteAgent/src/SimpleSocket/ipc/echo.sock")
	if err != nil {
		println("listen error", err)
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			println("accept error", err)
			return
		}

		go echoServer(fd)
	}
}
