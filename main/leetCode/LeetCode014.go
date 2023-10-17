package main

import "math"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	temp := strs[0]
	for i := 1; i < len(strs); i++ {
		temp = checkStrings(temp, strs[i])
		if len(temp) == 0 {
			break
		}
	}
	return temp
}

func checkStrings(temp string, s string) string {
	length := math.Min(float64(len(temp)), float64(len(s)))
	index := 0
	for index < int(length) && temp[index] == s[index] {
		index++
	}
	return temp[:index]
}
