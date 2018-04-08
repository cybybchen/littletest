package main

import (
	"fmt"
	"reflect"
)

type testinterface interface {
	Start()
}

type test struct {
	a uint32
}

func (this *test) Start() {
	this.a = 1
}

func (this *test) do() {
	this.a = 1
}

func main() {
	//go func() {
	//	fmt.Println("aaa")
	//}()
	//time.Sleep(time.Second)
	//ss := []int{1, 2, 3}
	//for i := 4; i < 10; i++ {
	//	fmt.Printf("addr:%p,len:%d,cap:%d\n", ss, len(ss), cap(ss))
	//	ss = append(ss, i)
	//}
	//a := 2
	//a++
	//fmt.Println(a)
	//var t *test
	//t1 := &test{}
	//t1.a = 10
	//t2 := new(test)
	//fmt.Println(t)
	//fmt.Println(t1)
	//fmt.Println(t2)

	//t := &test{}
	//testif(t)
	//data := []int{0, 1, 2, 3, 4, 5}
	//s := data[:3]
	//fmt.Println(s)
	//a := []test{}
	//t := test{}
	//t.a = 10
	//a = append(a, t)
	//t.a = 20
	//a = append(a, t)
	//fmt.Println(a[0])
	//fmt.Println(a[1])
	//a := "haha" + strconv.FormatUint(10, 10) + "haha"
	//fmt.Println(a)

	//errmap := make(map[uint32][]uint32)
	////errmap[1] = append(errmap[1], 1)
	////errmap[2] = append(errmap[2], 2)
	//fmt.Println(errmap)
	//fmt.Println(len(errmap))

	a := []uint32{1, 2, 3, 4, 5, 6, 7}
	b := a[:]
	b[0] = 10
	fmt.Println(a)
	fmt.Println(b)
}

func testif (t testinterface) {
	p := reflect.ValueOf(t)
	to := reflect.TypeOf(t)
	fmt.Println(to)
	fmt.Println(p)
	fmt.Println(p.Type())
	fmt.Println(p.Kind())
	fmt.Println(p.Elem())
	v := p.Elem()
	fmt.Println(v.Kind())
	//namef := v.FieldByName("start")
	//fmt.Println(namef)
	//names := namef.String()
	//fmt.Println(names)
	pt := p.Type()
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("11")
		mt := pt.Method(i)
		fmt.Println(mt.Name)
	}

	//str := "uuzuGtaGetuser"
	//
	//str2 := es.EscapeHandler(str)
	//fmt.Println(str2)
}