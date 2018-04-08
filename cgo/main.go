package main

/*
#include <stdio.h>
#include <stdlib.h>
void test(char *s) {
printf("%s\n", s);
}
char* cstr() {
return "abcde";
}
*/
import "C"
import (
	"unsafe"
	"fmt"
)

func main() {
	s := "Hello, World!"
	cs := C.CString(s) // 该函数在 C heap 分配内存，需要调⽤用 free 释放。
	defer C.free(unsafe.Pointer(cs)) // #include <stdlib.h>
	C.test(cs)
	cs = C.cstr()
	fmt.Println(C.GoString(cs))
	fmt.Println(C.GoStringN(cs, 2))
	fmt.Println(C.GoBytes(unsafe.Pointer(cs), 2))
}
