package main

import "fmt"

/*
找出字符串中第一个匹配项的下标
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。
输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
*/
// 暴力破解
func strStr(haystack string, needle string) int {
	i, j := 0, 0
	for j < len(haystack) {
		if len(haystack)-j < len(needle) {
			return -1
		}
		k := j
		for i < len(needle) {
			if haystack[k] != needle[i] {
				i = 0
				j++
				break
			}
			i++
			k++
		}
		if i == len(needle) {
			return k - len(needle)
		}
	}
	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issipi"
	fmt.Println(strStr(haystack, needle))
}
