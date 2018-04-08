package main

import (
	"fmt"
	"littletest/rpc/model"
	"log"
	"net/rpc"
)

func main() {
	serverAddress := ""
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &model.Args{7,8}
	var reply = &model.Reply{}
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Asynchronous call
	reply1 := new(model.Reply)
	divCall := client.Go("Arith.Div", args, reply1, nil)
	fmt.Println(reply1)
	divCall = <-divCall.Done	// will be equal to divCall
	// check errors, print, etc.
	fmt.Println(reply1)
}
