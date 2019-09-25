package main

import (
	"fmt"
	"io/ioutil"
)

const filename string = "abc.txt"

//	if 语句
func readFile() {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func readFile2() {
	// if 还可以类似 for 使用
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// switch 语句
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}

	return result
}

func main() {
	//readFile()
	//readFile2()

	fmt.Println(eval(1, 20, "+"))
}
