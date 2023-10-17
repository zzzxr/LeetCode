package main

import (
	"fmt"
)

// 有一个字符串数组，每个字符串都只包含小写字母，现在需要找到两个长度相乘最大的字符串，并且两个字符串不能有相同的字母，如果没有满足这个条件的结果，
//返回0
//
// Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
// Output: 16
//
// Input: ["a","aa","aaa","aaaa"]
// Output: 0

func main() {
	s := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(s)
	fmt.Println(test04(s))
}

func test04(s []string) int {
	if len(s) == 0 {
		return 0
	}
	maxLength := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if !swapString(s[i], s[j]) {
				continue
			}
			maxLength = swapMax(maxLength, len(s[i])*len(s[j]))
		}
	}
	return maxLength
}

func swapString(a, b string) bool {
	stringMap := map[uint8]int{}
	for i := 0; i < len(a); i++ {
		stringMap[a[i]] = 1
	}
	for i := 0; i < len(b); i++ {
		if _, ok := stringMap[b[i]]; ok {
			return false
		}
	}
	return true
}

func swapMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
