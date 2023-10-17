package main

import "fmt"

/*
切片是值传递还是引用传递
*/
func main() {
	a := make([]int, 10, 10)
	fmt.Println("a", a)
	b := a
	b[0] = 10
	fmt.Println(a, b)
}
