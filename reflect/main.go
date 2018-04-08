package main

import (
	"fmt"
	"reflect"
)

//type User struct {
//	Name string `field:"username" type:"nvarchar(20)"`
//	Age int `field:"age" type:"tinyint"`
//}
//func main() {
//	var u User
//	t := reflect.TypeOf(u)
//	f, _ := t.FieldByName("Name")
//	fmt.Println(f.Tag)
//	fmt.Println(f.Tag.Get("field"))
//	fmt.Println(f.Tag.Get("type"))
//}

//var (
//	Int = reflect.TypeOf(0)
//	String = reflect.TypeOf("")
//)
//func main() {
//	c := reflect.ChanOf(reflect.SendDir, String)
//	fmt.Println(c)
//	m := reflect.MapOf(String, Int)
//	fmt.Println(m)
//	s := reflect.SliceOf(Int)
//	fmt.Println(s)
//	t := struct{ Name string }{}
//	p := reflect.PtrTo(reflect.TypeOf(t))
//	fmt.Println(p)
//}

func main() {
	t := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t)
}
