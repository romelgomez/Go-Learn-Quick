package _func

import (
	"fmt"
	"reflect"
	"runtime"
)

// func 可以类似于元组返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

func usageDiv(a, b int) int {
	q, _ := div3(a, b)
	return q
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling func %s  with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func main() {
	//fmt.Println(div(13, 3))
	//q, r := div2(14, 4)
	//fmt.Println(q, r)
	fmt.Println(apply(usageDiv, 19, 4))
}
