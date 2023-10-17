package main

import "strconv"

/*
二进制求和
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
输入:a = "11", b = "1"
输出："100"
*/
func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	ans := ""
	carry := 0
	m := max(lenA, lenB)
	for i := 0; i < m; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0') // ASCII码转换
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
