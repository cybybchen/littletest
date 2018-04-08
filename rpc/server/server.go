package main

import (
	"littletest/rpc/model"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	arith := new(model.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	time.Sleep(time.Second * 100000)
}
