package array

import (
	"fmt"
)

type array struct {

}

func (this *array) handle() {
	fmt.Println(a)
}

var a = [3]int{1, 2} // 未初始化元素值为 0。
var b = [...]int{1, 2, 3, 4} // 通过初始化值确定数组⻓长度。
var c = [5]int{2: 100, 4:200} // 使⽤用索引号初始化元素。
var d = [...]struct {
name string
age uint8
}{
{"user1", 10}, // 可省略元素类型。
{"user2", 20}, // 别忘了最后⼀一⾏行的逗号。
}
