package main

import "fmt"

func main() {
	// basic usage about pointer
	//var a = 2
	//var pa = &a
	//*pa = 3
	//fmt.Println(a)
	// Go 指针不能够参与运算
	// Go 语言只有值传递
	a, b := 3, 4
	swap(a, b)
	// 换不过来
	fmt.Println(a, b)
	// 换过来
	swap2(&a, &b)
	fmt.Println(a, b)
	// 换过来
	a, b = swap3(a, b)
	fmt.Println(a, b)
}

func swap(a, b int) {
	b, a = a, b
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {
	return b, a
}
