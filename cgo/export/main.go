package main

import "fmt"

/*
#include "test.h"
*/
import "C"
//export hello
func hello() {
	fmt.Println("Hello, World!\n")
}
func main() {
	C.test()
}
