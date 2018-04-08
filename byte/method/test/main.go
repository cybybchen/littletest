package main

import "fmt"

type X struct{}
func (this *X) test() {
	println("X.test")
	fmt.Printf("Value: %p\n", this)
}
func main() {
	p := X{}
	p.test()
	p1 := &p
	p1.test()
	// Error: calling method with receiver &p (type **X) requires explicit dereference
	//(&p).test()
}
