package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//if len(os.Args) <= 1 {
	//	fmt.Println("usage: go run client2.go YOUR_CONTENT")
	//	return
	//}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8887")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	time.Sleep(time.Second * 2)
	//data := os.Args[1]
	//data := "hellohahashahaerarfawgawe"
	data := make([]byte, 65536)
	conn.Write(data)

	time.Sleep(time.Second * 10000)
}
