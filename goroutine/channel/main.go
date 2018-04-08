package main

import (
	A "fmt"
)

//func main() {
//	a, b := make(chan int, 3), make(chan int)
//	go func() {
//		v, ok, s := 0, false, ""
//		for {
//			select { // 随机选择可⽤用 channel，接收数据。
//			case v, ok = <-a: s = "a"
//			case v, ok = <-b: s = "b"
//			}
//			if ok {
//				fmt.Println(s, v)
//			} else {
//				os.Exit(0)
//			}
//		}
//	}()
//	for i := 0; i < 5; i++ {
//		select { // 随机选择可⽤用 channel，发送数据。
//		case a <- i:
//		case b <- i:
//		}
//	}
//	close(a)
//	fmt.Println("hahahaha")
//	select {} // 没有可⽤用 channel，阻塞 main goroutine。
//	fmt.Println("haha")
//}

//func main() {
//	w := make(chan bool)
//	c := make(chan int, 2)
//	go func() {
//		select {
//		case v := <-c: fmt.Println(v)
//		case <-time.After(time.Second * 3): fmt.Println("timeout.")
//		}
//		w <- true
//	}()
//	 //c <- 1 // 注释掉，引发 timeout。
//	<-w
//}

//c := make(chan int, 3)
//var send chan<- int = c // send-only
//var recv <-chan int = c // receive-only
//send <- 1
//// <-send // Error: receive from send-only type chan<- int
//<-recv
//// recv <- 2 // Error: send to receive-only type <-chan int
//不能将单向 channel 转换为普通 channel。
//d := (chan int)(send) // Error: cannot convert type chan<- int to type chan int
//d := (chan int)(recv)

type Request struct {
	data []int
	ret chan int
}
func NewRequest(data ...int) *Request {
	return &Request{ data, make(chan int, 1) }
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}
func main() {
	req := NewRequest(10, 20, 30)
	Process(req)
	//fmt.Println(<-req.ret)
	A.Println("111")
}