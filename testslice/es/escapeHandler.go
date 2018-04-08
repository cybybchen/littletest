package es

import "bytes"

func EscapeHandler(str string) string {
	var buff = []byte(str)
	var arr = make([][]byte, 0)
	var pre = 0
	for i, ch := range buff {
		if ch >= 'A' && ch < 'Z' {
			ch += 'a' - 'A'
			buff[i] = ch
			arr = append(arr, buff[pre:i])
			pre = i
		}
	}

	if pre != 0 {
		arr = append(arr, buff[pre:])
	}

	if len(arr) == 0 {
		return str
	}

	str = string(bytes.Join(arr, []byte("/")))
	return str
}