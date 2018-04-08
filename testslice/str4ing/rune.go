package main

import "fmt"

func main() {
	s := "ababa eee  haha"
	r := []rune(s)
	for _, v := range r {
		fmt.Println(string(v))
	}
	fmt.Println(r)
}
