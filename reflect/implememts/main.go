package main

import (
	"fmt"
	"reflect"
)

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var p *int
//	var x interface{} = p
//	fmt.Println(x == nil)
//	v := reflect.ValueOf(p)
//	fmt.Println(v.Kind(), v.IsNil())
//}

type User struct {
	Username string
	age int
}
func main() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)
	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
}
