package main

import (
	"fmt"
	"github.com/funny/pprof"
	_"unsafe"
)

type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

func main() {
	//t := &T{1, 2, 3, "", true}
	//fmt.Println(unsafe.Sizeof(*t))
	//fmt.Println(unsafe.Sizeof(t.t1))
	//fmt.Println(unsafe.Sizeof(t.t2))
	//fmt.Println(unsafe.Sizeof(t.t3))
	//fmt.Println(unsafe.Sizeof(t.t4))
	//fmt.Println(unsafe.Sizeof(t.t5))

	v := 300
	b := make([]byte, 2)
	b[0] = byte(v >> 8)
	b[1] = byte(v)

	fmt.Println(b)

	summary := pprof.GCSummary()

	println(summary.String())
}
