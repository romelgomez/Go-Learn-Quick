package container

import (
	"fmt"
)

func main() {
	// new map
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)

	var m3 map[string]string

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing Map")
	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	// get value
	fmt.Println("get value")
	courseName := m["course"]
	fmt.Println(courseName)
	fmt.Println(m["234"]) //不存在的 返回空

	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		panic("key does not exist")
	}

	// delete value
	delete(m, "course")
	fmt.Println(m)
}
