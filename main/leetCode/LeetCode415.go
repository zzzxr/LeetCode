package main

import "strconv"

/*
字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
输入：num1 = "11", num2 = "123"
输出："134"
*/
func addStrings(num1 string, num2 string) string {
	lenA, lenB := len(num1), len(num2)
	ans := ""
	carry := 0
	m := max(lenA, lenB)
	for i := 0; i < m; i++ {
		if i < lenA {
			carry += int(num1[lenA-i-1] - '0') // ASCII码转换
		}
		if i < lenB {
			carry += int(num2[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%10) + ans
		carry /= 10
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
