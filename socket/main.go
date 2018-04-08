package main

import (
	"fmt"
	"net"
	"time"
)

var port = ":8085"

func handleConnection() {
	conn, e := net.Dial("tcp", port)
	if e != nil {
		fmt.Println(e)
	}
	buf := make([]byte, 512)
	for {
		fmt.Println("11")
		n, e := conn.Read(buf)
		fmt.Println("22")
		if e == nil {
			fmt.Println(buf[:n])
		} else {
			fmt.Println(e)
		}
	}
}

func handleSend(conn net.Conn) {
	for {
		time.Sleep(time.Second * 5)
		conn.Write([]byte{1, 2})
	}
}

func main() {
	//addr, err := net.ResolveTCPAddr("tcp", ":80")
	//if err != nil {
	//	fmt.Println(err)
	//}
	listen, e := net.Listen("tcp", port)
	if e != nil {
		fmt.Println(e)
	}
	go handleConnection()
	for {
		fmt.Println("333")
		conn, err := listen.Accept()
		fmt.Println("444")
		if err != nil {
			fmt.Println(err)
		}
		go handleSend(conn)
	}
}
