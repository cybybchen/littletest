package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"reflect"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

type Test struct {
	Cyb uint32
	ycb uint32
}

func main() {
	//var str1, str2 string
	//str1 = getCurrentDirectory()
	//fmt.Println(str1)
	//str2 = getParentDirectory(str1)
	//fmt.Println(str2)

	str := "cyb"
	fmt.Println(str[len(str) - 1:])

	test := &Test{Cyb: 10, ycb: 5}
	rf := reflect.ValueOf(test)
	fmt.Println(rf.Kind())
	rv := rf.Elem()
	fmt.Println(rv.Kind())
	pt := rv.Type()
	fmt.Println(pt.Kind())
	for i := 0; i < pt.NumField(); i++ {
		//mt := rv.Field(i)
		fmt.Println(pt.Field(i).Name)
		fmt.Println(rv.Field(i).Uint())
	}
}