package runes

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var s = "Yes我爱好一阵!" // UTF-8
	fmt.Println(len(s))

	// 获得字节数组
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println("")
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println("")
	// 获得字符数量
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	// rune 相当于 Go 的 char
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
