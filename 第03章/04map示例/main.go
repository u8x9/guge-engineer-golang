package main

import (
	"fmt"
)

// 寻找最长不含有重复字符的子串
// 对于每个字母x：
// + last[x] 不存在，或者 < start ==> 无需操作
// + last[x] >= start ==> 更新 start
// + 更新 last[x], 更新 maxLength
func lengthOfNonRepeatingSubString(s string) int {
	last := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastIdx, ok := last[ch]; ok && lastIdx >= start {
			start = lastIdx + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubString("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubString("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubString("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubString(""))
	fmt.Println(lengthOfNonRepeatingSubString("b"))
	fmt.Println(lengthOfNonRepeatingSubString("abcdefg"))
	fmt.Println(lengthOfNonRepeatingSubString("锄禾日当午"))
	fmt.Println(lengthOfNonRepeatingSubString("一二三二一"))
}
