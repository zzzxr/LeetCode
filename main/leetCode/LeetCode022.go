package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 1
输出：["()"]
*/

var temp []string

func generateParenthesis(n int) []string {

	def(n, 0, 0, "")
	return temp
}

func def(n, lc, rc int, path string) {
	if lc == n && rc == n {
		temp = append(temp, path)
		return
	}
	if lc < n {
		def(n, lc+1, rc, path+"(")
	}
	if rc < lc {
		def(n, lc, rc+1, path+")")
	}
	return
}

func main() {
	temp = generateParenthesis(3)
	fmt.Println(temp)

}
