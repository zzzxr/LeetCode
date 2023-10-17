package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println(x) // 在延迟执行时，x 的值已经被捕获
	x++
	fmt.Println(x)
}
