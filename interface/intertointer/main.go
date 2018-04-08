package main

import "fmt"

type inter1 interface {
	do()
}

type inter2 interface {
	send()
	do()
}

type inter struct {

}

type test struct {
	in inter
}

type test1 struct {
	t test
}

func (this *inter) send() {

}

func (this *inter) do() {
	fmt.Println("inter")
}

func main() {
	//in := &inter{}
	//send2(in)
	test := &test1{}
	fmt.Println(test.t)
	//test.t.in.a = 10
	fmt.Println(test.t.in)
}

func send1(in1 inter1) {
	fmt.Println("1haha")
	in := in1.(*inter)
	in.do()
}

func send2(in2 inter2) {
	fmt.Println("2haha")
	send1(in2)
}
