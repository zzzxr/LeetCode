package main

import "fmt"

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
*/
func reverseString(s []byte) {
	var temp1 byte
	for i := 0; i < len(s); i++ {
		if i > len(s)/2 {
			break
		}
		temp1 = s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = temp1
	}
	fmt.Println(s)
}

func reverseString1(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	fmt.Println(s)
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString1(s)
}
