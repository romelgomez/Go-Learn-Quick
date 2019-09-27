package loop

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Go 里没有 while 语句
func convertToBin(n int) string {
	result := ""
	// for 可以省略初始条件，相当于 while
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	// 循环遍历
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13))
}
