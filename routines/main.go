package main

import (
	"fmt"
	"sync"
)

func main() {
	pt := fmt.Printf

	type SchedReq struct {
		C        chan struct{}
		Priority int
	}
	enters := make(chan SchedReq)
	reqs := make(chan SchedReq)
	quits := make(chan SchedReq)

	// scheduler
	const MaxPriority = 65535
	go func() {
		var hold [MaxPriority + 1][]SchedReq
		var nThreads [MaxPriority + 1]int
	loop:
		for {
			select {

			case req := <-enters:
				nThreads[req.Priority]++

			case req := <-reqs:
				for i := 0; i < req.Priority; i++ {
					if nThreads[i] > 0 {
						hold[req.Priority] = append(hold[req.Priority], req)
						continue loop
					}
				}
				req.C <- struct{}{}

			case req := <-quits:
				nThreads[req.Priority]--
				if nThreads[req.Priority] > 0 {
					continue loop
				}
				for i := req.Priority + 1; i <= MaxPriority; i++ {
					if len(hold[i]) > 0 {
						for _, req := range hold[i] {
							req.C <- struct{}{}
						}
						hold[i] = hold[i][0:0]
						break
					}
				}

			}
		}
	}()

	spawn := func(priority int, fn func(sched func())) {
		c := make(chan struct{})
		req := SchedReq{
			C:        c,
			Priority: priority,
		}
		sched := func() {
			reqs <- req
			<-c
		}
		enters <- req
		go func() {
			defer func() {
				quits <- req
			}()
			sched()
			fn(sched)
		}()
	}

	// demo
	start := make(chan struct{})
	wg := new(sync.WaitGroup)
	num := 10000
	repeat := 3
	wg.Add(num * repeat)

	for i := 0; i < num; i++ {
		i := i
		for range make([]struct{}, repeat) {
			spawn(i, func(sched func()) {
				// 确保同时进入 sched
				<-start
				sched()
				// 会按照优先级顺序执行下面这行
				pt("%d\n", i)
				wg.Done()
			})
		}
	}

	close(start)
	wg.Wait()

}
