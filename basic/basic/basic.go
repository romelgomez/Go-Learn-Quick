package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 可以在函数外定义变量
var aa = 3
var cc = "hhhdhfdf"
// 但是在函数外不能够使用 := 定义变量

// 不是全局变量，只是包内作用
var (
	dd = 4
	pp = true
)

func variableZeroValue() {
	// 变量值是在后面定义类型
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// 可以推断类型
	var a, b, c, d = 3, 4, true, "hhhh"
	var s = "abc"
	var char = 'a'
	fmt.Println(a, b, s, c, d, char)
}

func variableShorter() {
	// := 定义一个变量
	a, b, c, d := 3, 4, true, "hhhh"
	fmt.Println(a, b, c, d)
}

// 验证欧拉公式 e^i*pi + 1 = 0
func euler() {
	//c := 3 + 4i
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义
func consts() {
	// 常量在 Go 里面一般不会全部大写
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	fmt.Println(math.Sqrt(a*a+b*b), filename, c)
}

// 枚举类型定义
func enums() {

	// 自增值的
	const (
		cpp = iota
		_
		python
		golang
	)

	fmt.Println(cpp, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//variableShorter()
	//variableInitialValue()
	//variableZeroValue()
	//fmt.Println("Hello World")
	//fmt.Println(aa, cc)

	//euler()

	//triangle()

	//consts()

	//enums()
}
