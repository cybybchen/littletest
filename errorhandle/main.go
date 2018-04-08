package main

import (
	"fmt"
)

//func main() {
//	defer func() {
//		fmt.Println(recover())
//	}()
//	defer func() {
//		panic("defer panic")
//	}()
//	panic("test panic")
//}

func test() {
	defer recover() // ⽆无效！
	defer func() {
		fmt.Println(recover()) // ⽆无效！
	}()
	defer func() {
		func() {
			println("defer inner")
			recover() // ⽆无效！
		}()
	}()
	panic("test panic")
}

//func test1(x, y int) {
//	var z int
//	func() {
//		defer func() {
//			if recover() != nil {
//				z = 0
//				fmt.Println("haha")
//			}
//		}()
//		z = x / y
//		fmt.Println(z)
//		return
//	}()
//	println("x / y =", z)
//}
//
//func main() {
//	test1(20, 10)
//}
//var ErrDivByZero = errors.New("division by zero")
//
//func div(x, y int) (int, error) {
//	if y == 0 {
//		return 0, ErrDivByZero
//	}
//	return x / y, nil
//}
//func main() {
//	switch z, err := div(10, 0); err {
//	case nil:
//		println(z)
//	case ErrDivByZero:
//		//panic(err)
//		fmt.Println(err)
//	}
//}

func main() {
	//data := [...]int{0, 1, 2, 3, 4, 5, 6}
	//s2 := data[1:4:6]
	//slice := data[1:4:5]
	//s := slice[1:2]
	//fmt.Println(slice, len(slice), cap(slice))
	//fmt.Printf("%p", slice)
	//fmt.Printf("\n%p, %p", &s[0], s)
	//fmt.Printf("\n%p", s2)

	//s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使⽤用索引号。
	//fmt.Println(s1, len(s1), cap(s1))

	//d := [5]struct {
	//	x int
	//}{}
	//s := d[:]
	//d[1].x = 10
	//s[2].x = 20
	//fmt.Println(d)
	//fmt.Println(s)
	//fmt.Printf("%p, %p\n", &d, &d[0])

	//s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := s[2:5] // [2 3 4]
	//s2 := s1[2:6:7] // [4 5 6 7]
	//fmt.Printf(s1)
	//fmt.Println(s2)
	//s3 := s2[3:6] // Error
	//
	//fmt.Println(s3)

	//data := [...]int{0, 1, 2, 3, 4, 10: 0}
	//s := data[:2:3]
	//s = append(s, 100, 200) // ⼀一次 append 两个值，超出 s.cap 限制。
	//fmt.Println(s, data) // 重新分配底层数组，与原数组⽆无关。
	//fmt.Println(&s[0], &data[0])

	//s := make([]int, 0, 1)
	//c := cap(s)
	//for i := 0; i < 50; i++ {
	//	s = append(s, i)
	//	if n := cap(s); n > c {
	//		fmt.Printf("cap: %d -> %d\n", c, n)
	//		c = n
	//	}
	//}
}

