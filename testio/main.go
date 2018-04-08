package main

import (
	"fmt"
)

//func main() {
//		runtime.GOMAXPROCS(1)
//		int_chan := make(chan int, 1)
//		string_chan := make(chan string, 1)
//		int_chan <- 1
//		string_chan <- "Golang我们走，我们要做好朋友！！！"
//		//for {
//		select {
//		case value := <-int_chan:
//			fmt.Println(value)
//		case value := <-string_chan:
//			//panic(value) //总会随机到我，我会执行的。。。
//			fmt.Println(value)
//		default:
//			fmt.Println("break")
//		}
//		//}
//
//		select {
//		case value := <-int_chan:
//			fmt.Println(value)
//		case value := <-string_chan:
//			//panic(value) //总会随机到我，我会执行的。。。
//			fmt.Println(value)
//		default:
//			fmt.Println("break")
//		}
//}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func exeFibonacci(){
	c := make(chan int)
	quit := make(chan int)

	//produce data
	go func(){
		for i:= 0; i < 10; i++ {
			fmt.Println("channel data item ", <- c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func main() {
	//exeFibonacci()

	//fmt.Println(math.Trunc(1.634))
	//fmt.Println(math.Trunc(-1.234))
	//fmt.Println(math.Floor(1.634))
	//fmt.Println(math.Floor(-1.234))
	//fmt.Println(math.Ceil(1.234))
	//fmt.Println(math.Ceil(-1.234))

	//t := time.Now().Unix()
	//fmt.Println(t)
	//t1 := time.Now().UnixNano()
	//fmt.Println(t1)

	//b := []byte{10,11}
	//fmt.Println(b)
	//s := string(b)
	//fmt.Println(s)

	//s := "aaaa"
	//b := []byte(s)
	//fmt.Println(b)
	b := []byte{}
	b1 := make([]byte, 10)
	fmt.Println(b)
	fmt.Println(b1)
}