package main

import (
	"fmt"
)

func main() {
	data := make(chan int, 3) // 缓冲区可以存储 3 个元素
	exit := make(chan bool)
	data <- 1 // 在缓冲区未满前，不会阻塞。
	data <- 2
	data <- 3
	go func() {
		for d := range data { // 在缓冲区未空前，不会阻塞。
			fmt.Println(d)
		}
		exit <- true
	}()
	data <- 4 // 如果缓冲区已满，阻塞。
	data <- 5
	close(data)
	<-exit

	d1 := make(chan int)
	d2 := make(chan int, 3)
	d2 <- 1
	fmt.Println(len(d1), cap(d1)) // 0 0
	fmt.Println(len(d2), cap(d2))
}

//func main() {
//	wg := new(sync.WaitGroup)
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 6; i++ {
//			println(i)
//			if i == 3 { runtime.Gosched() }
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		stamp := time.Now().Unix()
//		time.Sleep(time.Second * 3)
//		delta := time.Now().Unix() - stamp
//		fmt.Println(delta)
//		println("Hello, World!")
//	}()
//	wg.Wait()
//}

//func main() {
//	wg := new(sync.WaitGroup)
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		defer println("A.defer")
//		func() {
//			defer println("B.defer")
//			runtime.Goexit() // 终⽌当前 goroutine
//			println("B")     // 不会执⾏
//		}()
//		println("A") // 不会执⾏
//	}()
//	wg.Wait()
//}
