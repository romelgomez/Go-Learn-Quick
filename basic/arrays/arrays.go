package arrays

import "fmt"

func main() {
	// 一维数组的定义
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	// 二维数组 (四行五列)
	var grid [4][5]int
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	// 下标, 值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//fmt.Println("----- Arr3 -----")
	//fmt.Println(arr3) // 第一个数字没有变成 100, 说明数组是值传递（会拷贝数组）
	//printArr3(arr3[:])
	//fmt.Println(arr3)

	// 对切片的修改会进入原数组，因为 Slice 是原数组的一个 View，包含了数组的指针、长度、容量
	//fmt.Println(arr3)

	//slices()

	appendSlice()
}

// 数组是值类型
func printArr(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}

	arr[0] = 100
}

func printArr2(arr *[5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}

	arr[0] = 100

	// 或者：
	(*arr)[0] = 101
}

// slice 类型是 []int, 与定长数组不同
func printArr3(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}

	arr[0] = 100
}

// Go 语言一般不直接使用数组，使用切片 Slice
func slices() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2:6] // 左开右闭区间 [2,3,4,5]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	s1 := arr[2:]

	// 对切片的修改会进入原数组，因为 Slice 是原数组的一个 View
	fmt.Println("After updateSlice(s1)")
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	// slice 可以向后扩展，不能向前扩展
	x1 := arr[2:6]
	x2 := x1[3:5]
	fmt.Println(x1, x2)
	fmt.Println(len(x1), cap(x1)) // 4, 6
	fmt.Println("----- Deleting Elements from Slice -----")
	si := append(arr[:3], arr[4:]...)
	fmt.Println(si)
}

// 添加元素
func appendSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s2 := arr[2:6]
	s3 := append(s2, 10) // arr = [0, 1, 2, 3, 4, 5, 6, 10]
	s4 := append(s3, 11) // arr = [0, 1, 2, 3, 4, 5, 6, 10, 11]
	s5 := append(s4, 12) // arr = [0, 1, 2, 3, 4, 5, 6, 10, 11]
	// 添加元素的时候如果超越了 capacity, 系统会自动分配一个更长的底层数组
	// 所以 s5 不再是 arr 的 view，修改只到 11
	fmt.Println("s3, s4, s5=", s3, s4, s5)
	fmt.Println("arr=", arr)
}

func updateSlice(s []int) {
	s[0] = 100
}
