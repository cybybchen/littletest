package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8887")
	if err != nil {
		log.Println("error listen:", err)
		return
	}
	defer l.Close()
	log.Println("listen ok")

	var i int
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		i++
		log.Printf("%d: accept a new connection\n", i)

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		time.Sleep(time.Second * 10)
		// read from the connection
		var buf = make([]byte, 65536)
		log.Println("start to read from conn")
		c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
		n, err := c.Read(buf)
		if err == io.EOF {
			log.Println("conn EOF read error:", err)
		}
		if err != nil {
			log.Println("conn read error:", err)
			//return
		}
		fmt.Println(len(buf))
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
