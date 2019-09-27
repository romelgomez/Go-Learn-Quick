package container

import "fmt"

// 寻找最长的不含重复字符的子串长度
// abcabcbb -> abc -> 3
// bbbbbb -> b -> b
// pwwkew -> wke -> 3
func main() {
	fmt.Println(solution("abcabcbb"))
}

func solution(input string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(input) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}


