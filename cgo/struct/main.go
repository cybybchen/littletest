package main

/*
#include <stdlib.h>
struct Data {
int x;
};
typedef struct {
int x;
} DataType;
struct Data* testData() {
return malloc(sizeof(struct Data));
}
DataType* testDataType() {
return malloc(sizeof(DataType));
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var d *C.struct_Data = C.testData()
	defer C.free(unsafe.Pointer(d))
	var dt *C.DataType = C.testDataType()
	defer C.free(unsafe.Pointer(dt))
	d.x = 100
	dt.x = 200
	fmt.Printf("%#v\n", d)
	fmt.Printf("%#v\n", dt)
}
