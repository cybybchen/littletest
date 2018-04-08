package main

import (
	_"fmt"
	"fmt"
)

func main() {
	//data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s := data[4:]
	//s2 := data[:5]
	//fmt.Println(s)
	//fmt.Println(s2)
	//copy(s2, s) // dst:s2, src:s
	//fmt.Println(s2)
	//fmt.Println(s)
	//fmt.Println(data)

	//m := map[string]int{
	//	"a": 1,
	//}
	//if v, ok := m["a"]; ok { // 判断 key 是否存在。
	//	println(v)
	//}
	//println(m["c"]) // 对于不存在的 key，直接返回 \0，不会出错。
	//m["b"] = 2 // 新增或修改。
	//delete(m, "c") // 删除。如果 key 不存在，不会出错。
	//println(len(m)) // 获取键值对数量。cap ⽆无效。
	//for k, v := range m { // 迭代，可仅返回 key。随机顺序返回，每次都不相同。
	//	println(k, v)
	//}
	//
	//fmt.Println(m)

	//type user struct{ name string }
	//m := map[int]user{ // 当 map 因扩张⽽而重新哈希时，各键值项存储位置都会发⽣生改变。 因此，map
	//	1: {"user1"}, // 被设计成 not addressable。 类似 m[1].name 这种期望透过原 value
	//} // 指针修改成员的⾏行为⾃自然会被禁⽌止。
	//m[1].name = "Tom"
	//type File struct {
	//	name string
	//	size int
	//	attr struct {
	//		perm int
	//		owner int
	//	}
	//}
	//f := File{
	//	name: "test.txt",
	//	size: 1025,
	//	//attr: {0755, 1}, // Error: missing type in composite literal
	//}
	//f.attr.owner = 1
	//f.attr.perm = 0755
	//var attr = struct {
	//	perm int
	//	owner int
	//}{2, 0755}
	//fmt.Println(attr.owner)
	//f.attr = attr
	//fmt.Println(f)

	//var u1 struct { name string "username" }
	//var u2 struct { name string }
	//u1 = u2 //can not

	type Resource struct {
		id int
		name string
	}
	type Classify struct {
		id int
	}
	type User struct {
		Resource // Resource.id 与 Classify.id 处于同⼀一层次。
		Classify
		name string // 遮蔽 Resource.name。
	}
	u := User{
		Resource{1, "people"},
		Classify{100},
		"Jack",
	}

	fmt.Println(u.Classify.id)
}



type User struct {
	name string
	age int
}



